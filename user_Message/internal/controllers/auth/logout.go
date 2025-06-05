package auth

import (
	"net/http"
	"web_userMessage/user_Message/internal/middleware"
)

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, _ := middleware.Store.Get(r, middleware.StoreName)
		// 清空 MySession 数据
		session.Values = make(map[interface{}]interface{})
		err := session.Save(r, w)
		if err != nil {
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
