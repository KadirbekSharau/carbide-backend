package controllers

import (
	"net/http"

	"github.com/KadirbekSharau/carbide-backend/src/dto"
	"github.com/KadirbekSharau/carbide-backend/src/services/auth"
	"github.com/KadirbekSharau/carbide-backend/src/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var roles = map[string]int{"seeker": 1, "employer": 2, "admin": 0} 
const expTime = 24*60*1

var AuthConfig = util.ErrorConfig{
	Options: map[string]util.ErrorMetaConfig{
		"Email required": {
			Tag:     "required",
			Field:   "Email",
			Message: "email is required on body",
		},
		"Email format not valid": {
			Tag:     "email",
			Field:   "Email",
			Message: "email format is not valid",
		},
		"Password required": {
			Tag:     "required",
			Field:   "Password",
			Message: "password is required on body",
		},
	},
}

type UserController interface {
	UserLogin(ctx *gin.Context)
	UserRegister(ctx *gin.Context)
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{service: service}
}

/* User Login Handler */
func (h *userController) UserLogin(ctx *gin.Context) {
	var input dto.InputLogin
	ctx.ShouldBindJSON(&input)
	if errResponse, errCount := util.GoValidator(&input, AuthConfig.Options); errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	data, status, err := h.service.UserLogin(&input)

	if status != 200 {
		util.APIResponse(ctx, err, status, http.MethodGet, data)
		return
	}
	accessTokenData := map[string]interface{}{"id": data.ID, "email": data.Email, "role": roles["user"]}
	h.createToken(accessTokenData, ctx, err)
}

/*  User Register Handler */
func (h *userController) UserRegister(ctx *gin.Context) {
	var input dto.InputUserSeekerRegister
	ctx.ShouldBindJSON(&input)
	if errResponse, errCount := util.GoValidator(&input, AuthConfig.Options); errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	data, status, err := h.service.UserRegister(&input)

	if status != 201 {
		util.APIResponse(ctx, err, status, http.MethodGet, data)
		return
	}
	accessTokenData := map[string]interface{}{"id": data.ID, "email": data.Email}
	h.createToken(accessTokenData, ctx, "Register new user account successfully")
}

func (h *userController) createToken(token map[string]interface{}, ctx *gin.Context, message string) {
	accessToken, errToken := util.Sign(token, "JWT_SECRET", expTime)
	if errToken != nil {
		defer logrus.Error(errToken.Error())
		util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}
	util.APIResponse(ctx, message, http.StatusCreated, http.MethodPost, map[string]string{"accessToken": accessToken})
} 