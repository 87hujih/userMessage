package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SendMessage 错误信息处理
func SendMessage(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(APIError{
		Code:    code,
		Message: message,
	})
	if err != nil {
		log.Println(err)
		return
	}
}
