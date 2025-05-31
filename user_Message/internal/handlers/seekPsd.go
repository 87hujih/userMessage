package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	md "web_userMessage/user_Message/internal/models"
	"web_userMessage/user_Message/pkg/utils"
)

// SeekPsd 忘记密码
func SeekPsd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		files, err := template.ParseFiles("user_Message/internal/views/user/seekPsd.html")
		if err != nil {
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
		phone := r.FormValue("phone")
		password := r.FormValue("password")
		err = md.ChangePsd(phone, password)
		if err != nil {
			if errors.Is(err, utils.ERROR_USER_NOTEXISTS) {
				utils.SendMessage(w, 400, "用户未注册")
			} else {
				log.Printf("数据库错误: %v", err)
				utils.SendMessage(w, 500, "内部服务器错误")
			}
			return
		}
		// 修改成功
		utils.SendMessage(w, 200, "密码已修改成功！")
	}
}
