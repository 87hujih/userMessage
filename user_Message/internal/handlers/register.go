package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	md "web_userMessage/user_Message/internal/models"
	"web_userMessage/user_Message/pkg/utils"
)

// Register 用户注册处理
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		files, err := template.ParseFiles("user_Message/internal/views/user/register.html")
		if err != nil {
			log.Println("注册页面展示失败", err)
			http.Error(w, "This page does not exist", http.StatusNotFound)
			return
		}
		err = files.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			return
		}
		name := r.FormValue("username")
		password := r.FormValue("password")
		phone := r.FormValue("phone")
		//注册账号到数据库
		err = md.RegisterUser(name, password, phone)
		if err != nil {
			if errors.Is(err, utils.ERROR_USER_EXISTS) {
				utils.SendMessage(w, 400, "手机号已被注册")
			} else {
				utils.SendMessage(w, 500, "注册失败")
			}
			return
		}
		// 注册成功
		utils.SendMessage(w, 200, "注册成功")
	}
}
