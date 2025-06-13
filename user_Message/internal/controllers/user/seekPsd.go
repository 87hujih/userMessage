package user

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"web_userMessage/user_Message/internal/service"
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

		//处理修改密码业务
		err = service.SeeKPsdService(phone, password)
		if err != nil {
			if errors.Is(err, utils.ERROR_USER_NOTEXISTS) {
				utils.SendMessage(w, 400, "该用户未注册")
			} else {
				log.Printf("%v", err)
				utils.SendMessage(w, 400, err.Error())
			}
			return
		}
		// 修改成功
		utils.SendMessage(w, 200, "密码已修改成功！")
	}
}
