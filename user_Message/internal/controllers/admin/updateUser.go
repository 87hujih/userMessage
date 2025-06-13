package admin

import (
	"log"
	"net/http"
	"strconv"
	"web_userMessage/user_Message/internal/service"
	"web_userMessage/user_Message/pkg/utils"
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

	//处理更新用户信息业务
	err = service.UpdateUserService(username, age, email, gender, int64(id))
	if err != nil {
		utils.SendMessage(w, 500, err.Error())
		log.Println(err)
		return
	}

	// 信息保存成功
	utils.SendMessage(w, 200, "修改成功")
}
