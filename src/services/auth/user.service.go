package services

import (
	"github.com/KadirbekSharau/carbide-backend/src/dto"
	model "github.com/KadirbekSharau/carbide-backend/src/models"
)

type UserService struct {
	repo *model.UserRepository
}

func NewUserService(repo *model.UserRepository) *UserService {
	return &UserService{repo: repo}
}

/* User Login Service */
func (s *UserService) UserLogin(input *dto.InputLogin) (*model.Users, int, string) {
	return s.repo.UserLogin(input)
}

/* User Registration Service */
func (s *UserService) UserRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string) {
	return s.repo.UserRegister(input)
}

/* Admin User Registration Service */
func (s *UserService) AdminRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string) {
	return s.repo.AdminRegister(input)
}