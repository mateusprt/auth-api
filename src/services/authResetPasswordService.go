package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mateusprt/auth-api/src/dtos"
	"github.com/mateusprt/auth-api/src/repositories"
	"gorm.io/gorm"
)

type ResetService struct {
	Repository *repositories.UsersRepository
}

func NewResetService(db *gorm.DB) *ResetService {
	repository := repositories.NewUserRepository(db)
	return &ResetService{Repository: repository}
}

func (service ResetService) Execute(data dtos.ResetDto) error {

	if data.Email == "" {
		return errors.New("Email can't be blank.")
	}

	userFound := service.Repository.FindUserByEmail(data.Email)

	if userFound.ID == 0 {
		return errors.New("User not found")
	}

	err := service.Repository.Update(&userFound)

	if err != nil {
		return errors.New("Unexpected error")
	}

	resetToken := generateResetToken()
	current_date := time.Now()

	userFound.ResetPasswordToken = resetToken
	userFound.ResetTokenSentAt = &current_date

	err = service.Repository.Update(&userFound)

	if err != nil {
		return errors.New("Unexpected error.")
	}

	return nil
}

func generateResetToken() string {
	return uuid.NewString()
}
