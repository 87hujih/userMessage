package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// SendMessage 错误信息处理
func SendMessage(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    code,
		"message": message,
	})
	if err != nil {
		log.Println(err)
		return
	}
}
