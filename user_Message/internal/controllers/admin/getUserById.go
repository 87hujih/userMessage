package admin

import (
	"encoding/json"
	"log"
	"net/http"
	md "web_userMessage/user_Message/internal/models"
	cm "web_userMessage/user_Message/pkg/utils"
)

// GetUserById 通过当前点的id来编辑用户信息
func GetUserById(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	phone := r.FormValue("phone")
	user, err := md.GetUser(phone)
	if err != nil {
		cm.SendMessage(w, 500, "获取用户失败")
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"user": map[string]interface{}{
			"UserId":   user.UserId.Int64,
			"UserName": user.UserName.String,
			"Age":      user.Age.Int64,
			"Email":    user.Email.String,
			"Gender":   user.Gender.String}})
	if err != nil {
		log.Println(err)
		return
	}
}
