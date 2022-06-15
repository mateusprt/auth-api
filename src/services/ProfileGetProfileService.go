package services

import (
	"errors"

	"github.com/mateusprt/auth-api/src/models"
	"github.com/mateusprt/auth-api/src/repositories"
	"gorm.io/gorm"
)

type GetProfileService struct {
	Repository *repositories.UsersRepository
}

func NewGetProfileService(db *gorm.DB) *GetProfileService {
	repository := repositories.NewUserRepository(db)
	return &GetProfileService{Repository: repository}
}

func (service GetProfileService) Execute(id int) (models.User, error) {

	userFound := service.Repository.FindUserById(id)

	if userFound.ID == 0 {
		return userFound, errors.New("User not found")
	}

	return userFound, nil
}
