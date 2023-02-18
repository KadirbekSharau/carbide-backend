package db

import (
	"log"

	"github.com/KadirbekSharau/carbide-backend/src/dto"
	"github.com/KadirbekSharau/carbide-backend/src/models"
	authService "github.com/KadirbekSharau/carbide-backend/src/services/auth"
	"gorm.io/gorm"
)

func accountsDataMigrator(db *gorm.DB) *models.Users {
	registerRepository := models.NewUserRepository(db)
	registerService := authService.NewUserService(registerRepository)
	admin := dto.InputUserSeekerRegister{
		FullName:  "Admin1",
		Email:      "admin1@gmail.com",
		Password:   "admin532",
	}
	data, status, err := registerService.AdminRegister(&admin)
	if status != 201 {
		log.Println(err)
	}
	return data
}