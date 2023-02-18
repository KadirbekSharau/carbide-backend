package app

import (
	"github.com/KadirbekSharau/carbide-backend/src/config/db"
	"github.com/KadirbekSharau/carbide-backend/src/routes"
	"github.com/KadirbekSharau/carbide-backend/src/util"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
	server *gin.Engine
}

func New() *App {
	return &App{}
}

func (app *App) Init() {
	app.db = db.NewDatabaseConnection()
	app.server = gin.Default()

	if util.GodotEnv("GO_ENV") == "debug" {
		gin.SetMode(gin.DebugMode)
	} 
	if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} 
	if util.GodotEnv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	app.server.Use(
		cors.New(cors.Config{
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"*"},
			AllowWildcard: true,
		}),
	)

	app.server.Use(helmet.Default())
	app.server.Use(gzip.Gzip(gzip.BestCompression))

	routes.InitAuthRoutes(app.db, app.server)

	app.server.Run(":" + util.GodotEnv("GO_PORT"))
}