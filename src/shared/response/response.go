package response

import (
	"encoding/json"
	"net/http"

	dtos "github.com/mateusprt/auth-api/src/dtos"
)

func JSON(w http.ResponseWriter, statusCode int, hasError bool, data interface{}) {
	w.WriteHeader(statusCode)
	data = dtos.ResponseDto{
		Error: hasError,
		Data:  data,
	}
	_ = json.NewEncoder(w).Encode(data)
}
