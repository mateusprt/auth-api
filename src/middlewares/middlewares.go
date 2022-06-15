package middlewares

import (
	"net/http"

	"github.com/mateusprt/auth-api/src/services/security"
	"github.com/mateusprt/auth-api/src/shared/response"
)

func Autentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := security.ValidateToken(r)

		if err != nil {
			message := struct{ Message string }{Message: err.Error()}
			response.JSON(w, http.StatusUnauthorized, true, message)
			return
		}
		next(w, r)
	}
}
