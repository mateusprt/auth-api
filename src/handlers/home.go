package handlers

import (
	"fmt"
	"net/http"

	dtos "github.com/mateusprt/auth-api/src/dtos"
	services "github.com/mateusprt/auth-api/src/services"
	"github.com/mateusprt/auth-api/src/services/security"
	"github.com/mateusprt/auth-api/src/shared/response"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	userOnToken, err := security.ExtractUserId(r)

	if err != nil {
		response.JSON(w, http.StatusUnauthorized, true, err.Error())
		return
	}

	getProfileService := services.NewGetProfileService(db)

	userFound, err := getProfileService.Execute(userOnToken)

	if err != nil {
		message := dtos.MessageDto{
			Message: err.Error(),
		}
		response.JSON(w, http.StatusNotFound, true, message)
		return
	}

	message := struct{ Message string }{Message: fmt.Sprintf("Bem vindo, %v", userFound.Username)}

	response.JSON(w, http.StatusOK, false, message)
}
