package routes

import (
	"github.com/KadirbekSharau/carbide-backend/src/controllers"
	"github.com/KadirbekSharau/carbide-backend/src/models"
	"github.com/KadirbekSharau/carbide-backend/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Auth routes */
func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		repository = models.NewUserRepository(db)
		service    = services.NewUserService(repository)
		authHandler    = controllers.NewUserController(service)
	)

	groupRoute := route.Group("/api/v1/auth")
	groupRoute.POST("/user/login", authHandler.UserLogin)
	groupRoute.POST("/user/register", authHandler.UserRegister)
}
