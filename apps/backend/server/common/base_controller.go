package common

import (
	"encoding/json"
	"net/http"
)

type BaseController struct{}

func (bc *BaseController) Response(w http.ResponseWriter, data interface{}, statusCode int, message string) {
	response := map[string]interface{}{
		"data":    data,
		"message": message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
