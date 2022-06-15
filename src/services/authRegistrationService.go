package services

import (
	"errors"
	"log"
	"time"

	"github.com/mateusprt/auth-api/src/dtos"
	"github.com/mateusprt/auth-api/src/models"
	"github.com/mateusprt/auth-api/src/repositories"
	"github.com/mateusprt/auth-api/src/services/security"
	"gorm.io/gorm"
)

type RegistrationService struct {
	Repository *repositories.UsersRepository
}

func NewRegistrationService(db *gorm.DB) *RegistrationService {
	repository := repositories.NewUserRepository(db)
	return &RegistrationService{Repository: repository}
}

func (registrationService RegistrationService) Execute(data dtos.RegistrationDto) error {

	if data.Password == "" {
		return errors.New("Password can't be blank.")
	}

	if data.PasswordConfirmation == "" {
		return errors.New("Password confirmation can't be blank.")
	}

	if !passwordHasTheMinimumLenght(data.Password) {
		return errors.New("Password is too short. Minimum is 8 characters.")
	}

	if passwordAndPasswordConfirmationMatch(data.Password, data.PasswordConfirmation) {
		passwordEncrypted := security.EncryptPassword(data.Password)
		confirmationToken := security.GenerateToken()
		datetimeNow := time.Now()

		user := models.User{
			Username:                data.Username,
			Email:                   data.Email,
			PasswordEncrypted:       passwordEncrypted,
			ConfirmationToken:       confirmationToken,
			ConfirmationTokenSentAt: &datetimeNow,
		}

		// envia email
		err := registrationService.Repository.Create(&user)

		if err != nil {
			log.Println(err)
			return errors.New("Account registration fails.")
		}

		return nil
	} else {
		return errors.New("Password and password confirmation doesn't match.")
	}
}

func passwordAndPasswordConfirmationMatch(password string, passwordConfirmation string) bool {
	return password == passwordConfirmation
}

func passwordHasTheMinimumLenght(password string) bool {
	return len(password) >= 8
}
