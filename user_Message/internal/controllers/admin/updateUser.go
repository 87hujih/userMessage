package admin

import (
	"log"
	"net/http"
	"strconv"
	md "web_userMessage/user_Message/internal/models"
	cm "web_userMessage/user_Message/pkg/utils"
)

// UpdateUser 编辑用户信息
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
		return
	}
	userId := r.FormValue("userId")
	username := r.FormValue("username")
	age := r.FormValue("age")
	email := r.FormValue("email")
	gender := r.FormValue("gender")
	id, _ := strconv.Atoi(userId)
	err = md.AlterInformation(username, age, email, gender, int64(id))
	if err != nil {
		cm.SendMessage(w, 500, "信息修改失败")
		log.Println(err)
		return
	}
	// 信息保存成功
	cm.SendMessage(w, 200, "修改成功")
}
