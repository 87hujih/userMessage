package user

import (
	"fmt"
	"html/template"
	"net/http"
	"web_userMessage/user_Message/internal/middleware"
	md "web_userMessage/user_Message/internal/models"
	"web_userMessage/user_Message/internal/service"
	"web_userMessage/user_Message/pkg/utils"
)

// PerCenter 个人中心
func PerCenter(w http.ResponseWriter, r *http.Request) {
	session, err := middleware.Store.Get(r, middleware.StoreName)
	if err != nil {
		http.Error(w, "会话获取失败", http.StatusInternalServerError)
		return
	}

	// 获取登录手机号
	phone, ok := session.Values["phone"].(string)
	if !ok {
		utils.SendMessage(w, 401, "未登录，请先登录")
		return
	}

	// 获取用户信息
	user, err := service.GetUserByIdService(phone)
	if err != nil {
		utils.SendMessage(w, 500, "获取用户失败")
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
		utils.SendMessage(w, 400, "表单解析失败")
		return
	}

	username := r.FormValue("username")
	age := r.FormValue("age")
	email := r.FormValue("email")
	gender := r.FormValue("gender")

	//处理更新用户信息业务
	err := service.UpdateUserService(username, age, email, gender, userId)
	if err != nil {
		utils.SendMessage(w, 500, "保存信息失败")
		fmt.Println("保存信息失败:", err)
		return
	}

	utils.SendMessage(w, 200, "保存成功")
}
