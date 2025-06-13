package auth

import (
	"html/template"
	"log"
	"net/http"
	"web_userMessage/user_Message/internal/service"
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

		//处理注册用户业务
		err = service.RegisterService(name, password, phone)
		if err != nil {
			log.Println(err)
			utils.SendMessage(w, 400, err.Error())
			return
		}

		// 注册成功
		utils.SendMessage(w, 200, "注册成功")
	}
}
