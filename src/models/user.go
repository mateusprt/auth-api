package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                      int        `gorm:"primaryKey" json:"id"`
	Username                string     `gorm:"type:varchar(100); unique; not null" json:"username"`
	Email                   string     `gorm:"type:varchar(100); unique; not null; index" json:"email"`
	PasswordEncrypted       string     `gorm:"type:varchar(255); not null" json:"password_encrypted"`
	ConfirmationToken       string     `gorm:"type:varchar(255); index" json:"confirmation_token"`
	ConfirmationTokenSentAt *time.Time `json:"confirmation_token_sent_at"`
	ConfirmedAt             *time.Time `json:"confirmed_at"`
	UnconfirmedEmail        bool       `gorm:"default:true" json:"unconfirmed_email"`
	ResetPasswordToken      string     `gorm:"index" json:"reset_password_token"`
	ResetTokenSentAt        *time.Time `json:"reset_token_sent_at"`
	gorm.Model
}

func (user User) ConfirmationTokenSentIsValid() bool {
	one_day := time.Hour * 24
	limit_date_for_account_confirmation := user.ConfirmationTokenSentAt.Add(one_day)
	return user.UnconfirmedEmail && limit_date_for_account_confirmation.After(time.Now())
}

func (user User) ResetPasswordTokenSentIsValid() bool {
	one_day := time.Hour * 24
	limit_date_for_reset_password := user.ResetTokenSentAt.Add(one_day)
	return limit_date_for_reset_password.After(time.Now())
}
