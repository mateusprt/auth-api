package dtos

type RegistrationDto struct {
	Username             string `json:"username"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type ConfirmationDto struct {
	Email string `json:"email"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetDto struct {
	Email string `json:"email"`
}

type ResetConfirmationDto struct {
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// Shared
type ResponseDto struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

type MessageDto struct {
	Message string `json:"message"`
}
