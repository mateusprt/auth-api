package handlers

import (
	"encoding/json"
	"net/http"

	dtos "github.com/mateusprt/auth-api/src/dtos"
	services "github.com/mateusprt/auth-api/src/services"
	"github.com/mateusprt/auth-api/src/shared/response"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var dataOfRequest dtos.RegistrationDto
	_ = json.NewDecoder(r.Body).Decode(&dataOfRequest)

	registrationService := services.NewRegistrationService(db)

	err := registrationService.Execute(dataOfRequest)

	if err != nil {
		message := dtos.MessageDto{
			Message: err.Error(),
		}
		response.JSON(w, http.StatusNotFound, true, message)
		return
	}

	message := dtos.MessageDto{
		Message: "Your account has been successfully created. We've sent a verification link to your email address.",
	}
	response.JSON(w, http.StatusOK, false, message)
}

func ConfirmationHandler(w http.ResponseWriter, r *http.Request) {
	var dataOfRequest dtos.ConfirmationDto
	_ = json.NewDecoder(r.Body).Decode(&dataOfRequest)

	confirmationService := services.NewConfirmationService(db)

	confirmationToken := r.URL.Query().Get("confirmation_token")
	err := confirmationService.Execute(dataOfRequest, confirmationToken)

	if err != nil {
		message := dtos.MessageDto{
			Message: err.Error(),
		}
		response.JSON(w, http.StatusNotFound, true, message)
		return
	}

	message := dtos.MessageDto{
		Message: "Your email address was successfully verified.",
	}
	response.JSON(w, http.StatusOK, false, message)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var dataOfRequest dtos.LoginDto
	_ = json.NewDecoder(r.Body).Decode(&dataOfRequest)

	loginService := services.NewLoginService(db)

	jwt, err := loginService.Execute(dataOfRequest)

	if err != nil {
		message := dtos.MessageDto{
			Message: err.Error(),
		}
		response.JSON(w, http.StatusNotFound, true, message)
		return
	}

	response.JSON(w, http.StatusOK, false, struct{ Token string }{Token: jwt})
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	var dataOfRequest dtos.ResetDto
	_ = json.NewDecoder(r.Body).Decode(&dataOfRequest)

	resetService := services.NewResetService(db)

	err := resetService.Execute(dataOfRequest)

	if err != nil {
		message := dtos.MessageDto{
			Message: err.Error(),
		}
		response.JSON(w, http.StatusNotFound, true, message)
		return
	}

	message := dtos.MessageDto{
		Message: "We've sent a link to reset your password.",
	}
	response.JSON(w, http.StatusOK, false, message)
}

func ResetConfirmationHandler(w http.ResponseWriter, r *http.Request) {
	var dataOfRequest dtos.ResetConfirmationDto
	_ = json.NewDecoder(r.Body).Decode(&dataOfRequest)

	resetConfirmationService := services.NewResetConfirmationService(db)

	resetPasswordToken := r.URL.Query().Get("reset_password_token")
	err := resetConfirmationService.Execute(dataOfRequest, resetPasswordToken)

	if err != nil {
		message := dtos.MessageDto{
			Message: err.Error(),
		}
		response.JSON(w, http.StatusNotFound, true, message)
		return
	}

	message := dtos.MessageDto{
		Message: "Password changed.",
	}
	response.JSON(w, http.StatusOK, false, message)
}
