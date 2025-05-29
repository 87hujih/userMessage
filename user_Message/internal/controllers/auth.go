package controllers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
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
		session, err := Store.Get(r, StoreName)
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

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, _ := Store.Get(r, StoreName)
		// 清空 session 数据
		session.Values = make(map[interface{}]interface{})
		err := session.Save(r, w)
		if err != nil {
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
