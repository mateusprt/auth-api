package services

import (
	"errors"

	"github.com/mateusprt/auth-api/src/dtos"
	"github.com/mateusprt/auth-api/src/repositories"
	"github.com/mateusprt/auth-api/src/services/security"
	"gorm.io/gorm"
)

type LoginService struct {
	Repository *repositories.UsersRepository
}

func NewLoginService(db *gorm.DB) *LoginService {
	repository := repositories.NewUserRepository(db)
	return &LoginService{Repository: repository}
}

func (service LoginService) Execute(data dtos.LoginDto) (string, error) {

	if data.Email == "" {
		return "", errors.New("Email can't be blank.")
	}

	if data.Password == "" {
		return "", errors.New("Password can't be blank.")
	}

	userFound := service.Repository.FindUserByEmail(data.Email)

	if userFound.ID == 0 {
		return "", errors.New("The email or password is incorrect.")
	}

	if userFound.UnconfirmedEmail {
		return "", errors.New("The email or password is incorrect.")
	}

	passwordDoesntMatch := !security.PasswordMatch(userFound.PasswordEncrypted, data.Password)

	if passwordDoesntMatch {
		return "", errors.New("The email or password is incorrect.")
	}

	jwt := security.GenerateJWT(userFound)

	return jwt, nil
}
