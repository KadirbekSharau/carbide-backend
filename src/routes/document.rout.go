package routes

import (
	"cloud.google.com/go/storage"
	"github.com/KadirbekSharau/carbide-backend/src/controllers"
	"github.com/KadirbekSharau/carbide-backend/src/models"
	"github.com/KadirbekSharau/carbide-backend/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Auth routes */
func InitDocumentRoutes(client *storage.Client, db *gorm.DB, route *gin.Engine) {
	var (
		repository      = models.NewDocumentRepository(db)
		service         = services.NewDocumentService(repository, client)
		documentHandler = controllers.NewDocumentController(service)
	)

	groupRoute := route.Group("/api/v1/document")
	groupRoute.POST("/", documentHandler.CreateDocument)
	groupRoute.GET("/:id", documentHandler.GetDocumentByUserId)
}
