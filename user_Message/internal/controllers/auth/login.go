package auth

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"web_userMessage/user_Message/internal/controllers/user"
	md "web_userMessage/user_Message/internal/models"
	"web_userMessage/user_Message/pkg/utils"
)

// Login 用户登录处理
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		files, err := template.ParseFiles("user_Message/internal/views/user/login.html")
		if err != nil {
			log.Println(err)
			return
		}
		err = files.Execute(w, nil)
		return
	}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			return
		}
		phone := r.FormValue("phone")
		password := r.FormValue("password")
		err = md.LoginUser(phone, password)
		if err != nil {
			if errors.Is(err, utils.ERROR_USER_INFORMATION) {
				utils.SendMessage(w, 400, "账号或密码错误")
			} else {
				utils.SendMessage(w, 500, "内部服务器错误")
			}
			return
		}
		// 登录成功 设置session
		session, err := user.Store.Get(r, user.StoreName)
		if err != nil {
			log.Println(err)
			return
		}
		u, err := md.GetUser(phone)
		if err != nil {
			utils.SendMessage(w, 500, "获取用户失败")
		}
		session.Values["phone"] = phone
		session.Values["userId"] = u.UserId.Int64
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			return
		}
		utils.SendMessage(w, 200, "登录成功")
	}
}
