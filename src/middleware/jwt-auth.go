package middleware

import (
	"net/http"

	"github.com/KadirbekSharau/carbide-backend/src/util"
	"github.com/gin-gonic/gin"
)

type UnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func Auth(validRoles []int) gin.HandlerFunc {

	return gin.HandlerFunc(func(ctx *gin.Context) {

		var errResponses = []UnathorizatedError{
			{
				Status: "Forbidden",
				Code: http.StatusForbidden,
				Method: ctx.Request.Method,
				Message: "Authorization is required for this endpoint",
			},
			{
				Status: "Unathorizated",
				Code: http.StatusUnauthorized,
				Method: ctx.Request.Method,
				Message: "accessToken invalid or expired",
			},
			{
				Status: "Unathorizated",
				Code: http.StatusUnauthorized,
				Method: ctx.Request.Method,
				Message: "accessToken and Role invalid or expired",
			},
		}

		if ctx.GetHeader("Authorization") == "" {
			ctx.JSON(http.StatusForbidden, errResponses[0])
			defer ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		token, err := util.VerifyTokenHeader(ctx, "JWT_SECRET")

		rolesVal := util.DecodeToken(token)

		roleExists := false
		for validRole := range validRoles {
			if rolesVal.Claims.Role == validRole {
				roleExists = true
			}
		}

		if !roleExists {
			ctx.JSON(http.StatusUnauthorized, errResponses[2])
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errResponses[1])
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			// global value result
			ctx.Set("user", token.Claims)
			// return to next method if token is exist
			ctx.Next()
		}
	})
}
