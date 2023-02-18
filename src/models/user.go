package models

import (
	"github.com/KadirbekSharau/carbide-backend/src/dto"
	"net/http"

	"github.com/KadirbekSharau/carbide-backend/src/util"
	"gorm.io/gorm"
)

// Users is database entity for user
type Users struct {
	gorm.Model
	Email     string     `gorm:"type:varchar(50);unique;not null"`
	FullName  string     `gorm:"type:varchar(50)"`
	Password  string     `gorm:"type:varchar(255)"`
	Role 	  string
	Documents []Document `gorm:"foreignKey:UserID"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

/* User Login Repository Service */
func (r *UserRepository) UserLogin(input *dto.InputLogin) (*Users, int, string) {
	var users Users
	db := r.db.Model(&users)
	users.Email = input.Email
	users.Password = input.Password

	if db.Debug().Select("*").Where("email = ?", input.Email).Find(&users).RowsAffected < 1 {
		return &users, http.StatusNotFound, "User account is not registered"
	}
	if util.ComparePassword(users.Password, input.Password) != nil {
		return &users, http.StatusForbidden, "Password is wrong"
	}
	return &users, http.StatusOK, "Logged in successfully"
}

/* User Seeker Registration Repository */
func (r *UserRepository) UserRegister(input *dto.InputUserSeekerRegister) (*Users, int, string) {
	var users Users
	db := r.db.Model(&users)
	if db.Debug().Select("*").Where("email = ?", input.Email).Find(&users).RowsAffected > 0 {
		return &users, http.StatusConflict, "Email already exists"
	}
	users.Email = input.Email
	users.FullName = input.FullName
	users.Password = input.Password
	users.Role = "User"
	if db.Debug().Create(&users).Error != nil {
		return nil, http.StatusForbidden, "Registering new account failed"
	}
	db.Commit()
	return &users, http.StatusCreated, "Registered successfully"
}

/* Admin Registration Repository */
func (r *UserRepository) AdminRegister(input *dto.InputUserSeekerRegister) (*Users, int, string) {
	var users Users
	db := r.db.Model(&users)

	if db.Debug().Select("*").Where("email = ?", input.Email).Find(&users).RowsAffected > 0 {
		return &users, http.StatusConflict, "Email already exists"
	}
	users.Email = input.Email
	users.FullName = input.FullName
	users.Password = input.Password
	users.Role = "Admin"

	if db.Debug().Create(&users).Error != nil {
		return nil, http.StatusForbidden, "Registering new account failed"
	}
	db.Commit()
	return &users, http.StatusCreated, "Registered successfully"
}

func (entity *Users) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	return nil
}