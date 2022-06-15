package services

import (
	"errors"
	"time"

	"github.com/mateusprt/auth-api/src/dtos"
	"github.com/mateusprt/auth-api/src/repositories"
	"gorm.io/gorm"
)

type ConfirmationService struct {
	Repository *repositories.UsersRepository
}

func NewConfirmationService(db *gorm.DB) *ConfirmationService {
	repository := repositories.NewUserRepository(db)
	return &ConfirmationService{Repository: repository}
}

func (service ConfirmationService) Execute(data dtos.ConfirmationDto, token string) error {
	if token == "" {
		return errors.New("Token not found.")
	}

	userFound := service.Repository.FindUserByConfirmationToken(token)

	if userFound.ID == 0 {
		return errors.New("User not found.")
	}

	if !userFound.ConfirmationTokenSentIsValid() {
		return errors.New("Token invalid.")
	}

	time_now := time.Now()
	userFound.ConfirmationToken = ""
	userFound.ConfirmationTokenSentAt = nil
	userFound.UnconfirmedEmail = false
	userFound.ConfirmedAt = &time_now

	err := service.Repository.Update(&userFound)

	if err != nil {
		return errors.New("Unexpected error.")
	}

	return nil
}
