package auth

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	ms "web_userMessage/user_Message/internal/MySession"
	md "web_userMessage/user_Message/internal/models"
	"web_userMessage/user_Message/pkg/utils"
)

// handleGet 为Get请求，渲染页面
func handleGet(w http.ResponseWriter) error {
	files, err := template.ParseFiles("user_Message/internal/views/user/login.html")
	if err != nil {
		log.Println(err)
		return err
	}
	return files.Execute(w, nil)
}

// handlePost 处理 POST 请求的登录表单提交。
func handlePost(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		log.Println("表单解析失败", err)
		return err
	}
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	if err := md.LoginUser(phone, password); err != nil {
		if errors.Is(err, utils.ERROR_USER_INFORMATION) {
			utils.SendMessage(w, 400, "账号或密码错误")
		} else {
			utils.SendMessage(w, 500, "内部服务器错误")
		}
		return err
	}
	//通过用户手机号获取用户信息
	u, err := md.GetUser(phone)
	if err != nil {
		utils.SendMessage(w, 500, "获取用户失败")
		return err
	}
	// 登录成功 设置session
	err = ms.SetupSession(w, r, ms.Store, ms.StoreName, phone, u.UserId.Int64)
	if err != nil {
		log.Println("设置 session 失败:", err)
		utils.SendMessage(w, 500, "设置 session 失败")
		return err
	}
	utils.SendMessage(w, 200, "登录成功")
	return nil
}

// Login 用户登录处理
func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == "GET" {
		err = handleGet(w)
	} else if r.Method == "POST" {
		err = handlePost(w, r)
	} else {
		utils.SendMessage(w, 405, "请求方法错误")
	}
	if err != nil {
		log.Println("处理请求失败", err)
		utils.SendMessage(w, 500, "内部服务器错误")
	}
}
