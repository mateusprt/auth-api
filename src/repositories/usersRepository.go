package repositories

import (
	"github.com/mateusprt/auth-api/src/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{DB: db}
}

func (repo UsersRepository) Create(user *models.User) error {
	result := repo.DB.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo UsersRepository) FindUserByConfirmationToken(token string) models.User {
	var user models.User
	repo.DB.First(&user, "confirmation_token = ?", token)
	return user
}

func (repo UsersRepository) FindUserByEmail(email string) models.User {
	var user models.User
	repo.DB.First(&user, "email = ?", email)
	return user
}

func (repo UsersRepository) FindUserById(id int) models.User {
	var user models.User
	repo.DB.First(&user, "id = ?", id)
	return user
}

func (repo UsersRepository) FindUserByPasswordResetToken(token string) models.User {
	var user models.User
	repo.DB.First(&user, "reset_password_token = ?", token)
	return user
}

func (repo UsersRepository) Update(user *models.User) error {
	result := repo.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
