package admin

import (
	"encoding/json"
	"log"
	"net/http"
	"web_userMessage/user_Message/internal/service"
	"web_userMessage/user_Message/pkg/utils"
)

// GetUserById 通过当前点的id来编辑用户信息
func GetUserById(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	phone := r.FormValue("phone")

	//通过用户id获得用户信息
	user, err := service.GetUserByIdService(phone)
	if err != nil {
		utils.SendMessage(w, 500, "获取用户失败")
		return
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
