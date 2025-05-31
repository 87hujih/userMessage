package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	md "web_userMessage/user_Message/internal/models"
	cm "web_userMessage/user_Message/pkg/utils"
)

// PerCenter 个人中心
func PerCenter(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, StoreName)
	if err != nil {
		http.Error(w, "会话获取失败", http.StatusInternalServerError)
		return
	}

	// 获取登录手机号
	phone, ok := session.Values["phone"].(string)
	if !ok {
		cm.SendMessage(w, 401, "未登录，请先登录")
		return
	}

	// 获取用户信息
	user, err := md.GetUser(phone)
	if err != nil {
		cm.SendMessage(w, 500, "获取用户失败")
		return
	}

	switch r.Method {
	case "GET":
		handleGet(w, r, user)
	case "POST":
		userId, _ := session.Values["userId"].(int64)
		handlePost(w, r, userId)
	default:
		http.Error(w, "不支持的请求方法", http.StatusMethodNotAllowed)
	}
}

// 处理 PerCenter中GET 请求：渲染页面
func handleGet(w http.ResponseWriter, r *http.Request, user *md.User) {
	files, err := template.ParseFiles("user_Message/internal/views/personalCenter.html")
	if err != nil {
		http.Error(w, "页面不存在", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Username":  user.UserName.String,
		"AvatarURL": user.AvatarURL.String,
		"IsAdmin":   user.IsAdmin.Int64,
		"Age":       user.Age.Int64,
		"Email":     user.Email.String,
		"Gender":    user.Gender.String,
	}

	if err := files.Execute(w, data); err != nil {
		fmt.Println("模板执行失败:", err)
	}
}

// 处理 PerCenter中POST 请求：保存用户信息
func handlePost(w http.ResponseWriter, r *http.Request, userId int64) {
	if err := r.ParseForm(); err != nil {
		cm.SendMessage(w, 400, "表单解析失败")
		return
	}

	username := r.FormValue("username")
	age := r.FormValue("age")
	email := r.FormValue("email")
	gender := r.FormValue("gender")

	err := md.AlterInformation(username, age, email, gender, userId)
	if err != nil {
		cm.SendMessage(w, 500, "保存信息失败")
		fmt.Println("保存信息失败:", err)
		return
	}

	cm.SendMessage(w, 200, "保存成功")
}
