package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

// DeleterUser 删除用户
func DeleterUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	phone := r.FormValue("phone")
	session, err := Store.Get(r, StoreName)
	phoneNew, _ := session.Values["phone"].(string)
	if phone == phoneNew {
		cm.SendMessage(w, 500, "不能删除自己")
		log.Println(err)
		return
	}
	err = md.DeleteUserByPhone(phone)
	if err != nil {
		cm.SendMessage(w, 500, "删除用户失败")
		return
	}
	cm.SendMessage(w, 200, "删除成功")
}
