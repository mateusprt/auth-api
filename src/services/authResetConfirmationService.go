package services

import (
	"errors"

	"github.com/mateusprt/auth-api/src/dtos"
	"github.com/mateusprt/auth-api/src/repositories"
	"github.com/mateusprt/auth-api/src/services/security"
	"gorm.io/gorm"
)

type ResetConfirmationService struct {
	Repository *repositories.UsersRepository
}

func NewResetConfirmationService(db *gorm.DB) *ResetConfirmationService {
	repository := repositories.NewUserRepository(db)
	return &ResetConfirmationService{Repository: repository}
}

func (service ResetConfirmationService) Execute(data dtos.ResetConfirmationDto, token string) error {
	if token == "" {
		return errors.New("Token not found")
	}

	userFound := service.Repository.FindUserByPasswordResetToken(token)

	if userFound.ID == 0 {
		return errors.New("User not found.")
	}

	if !userFound.ResetPasswordTokenSentIsValid() {
		return errors.New("Token invalid.")
	}

	if data.Password == "" {
		return errors.New("Password can't be blank.")
	}

	if data.PasswordConfirmation == "" {
		return errors.New("Password confirmation can't be blank.")
	}

	if data.Password != data.PasswordConfirmation {
		return errors.New("Password doesn't match.")
	}

	newPasswordEncrypted := security.EncryptPassword(data.Password)
	userFound.PasswordEncrypted = newPasswordEncrypted
	userFound.ResetPasswordToken = ""
	userFound.ResetTokenSentAt = nil

	err := service.Repository.Update(&userFound)

	if err != nil {
		return errors.New("Unexpected error")
	}

	return nil
}
