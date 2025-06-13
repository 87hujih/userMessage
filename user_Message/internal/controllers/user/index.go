package user

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web_userMessage/user_Message/internal/middleware"
	"web_userMessage/user_Message/internal/service"
	cm "web_userMessage/user_Message/pkg/utils"
)

const (
	// AvatarDir 存放头像的目录
	AvatarDir = "user_Message/user_img"
	//登录目录
	loginName = "/login"
)

// Index 进入页面
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		files, err := template.ParseFiles("user_Message/internal/views/index.html")
		if err != nil {
			log.Panicln(err)
			return
		}
		err = files.Execute(w, nil)
		if err != nil {
			log.Panicln(err)
			return
		}
	}
}

// 自定义 FuncMap
func parseTemplateWithFuncs(path string, name string) (*template.Template, error) {
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"gt":  func(a, b int) bool { return a > b },
		"lt":  func(a, b int) bool { return a < b },
	}

	// 创建模板并注册函数
	t := template.New(name).Funcs(funcMap)

	// 解析模板文件
	files, err := t.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	// 检查是否成功加载模板
	if len(files.Templates()) == 0 {
		return nil, fmt.Errorf("模板为空，请检查模板内容和结构")
	}

	return files, nil
}

// HomePage 系统首页
func HomePage(w http.ResponseWriter, r *http.Request) {
	session, err := middleware.Store.Get(r, middleware.StoreName)
	if err != nil {
		http.Redirect(w, r, loginName, http.StatusFound)
		return
	}

	phone, ok := session.Values["phone"].(string)
	if !ok {
		http.Redirect(w, r, loginName, http.StatusFound)
		return
	}

	// 获取当前页码
	page := 1
	if p := r.FormValue("page"); p != "" {
		if pg, err := strconv.Atoi(p); err == nil && pg > 0 {
			page = pg
		}
	}
	limit := 10

	//处理页面所有用户展示业务
	err, allUser, user, total := service.HomePageService(phone, page, limit)

	if err != nil {
		cm.SendMessage(w, 500, err.Error())
	}
	totalPages := (total + limit - 1) / limit

	data := map[string]interface{}{
		"AllUser":     allUser,
		"Username":    user.UserName.String,
		"AvatarURL":   user.AvatarURL.String,
		"isAdmin":     user.IsAdmin.Int64 == 1,
		"UserCount":   total,
		"CurrentPage": page,
		"TotalPages":  totalPages,
	}

	// 加载模板
	files, err := parseTemplateWithFuncs("user_Message/internal/views/dashboard.html", "dashboard")
	if err != nil {
		http.Error(w, "页面不存在", http.StatusInternalServerError)
		return
	}

	// 执行指定模板
	err = files.ExecuteTemplate(w, "dashboard", data)
	if err != nil {
		fmt.Println("执行模板出错:", err)
		return
	}
}
