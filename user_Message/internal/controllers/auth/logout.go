package auth

import (
	"net/http"
	"web_userMessage/user_Message/internal/controllers/user"
)

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, _ := user.Store.Get(r, user.StoreName)
		// 清空 session 数据
		session.Values = make(map[interface{}]interface{})
		err := session.Save(r, w)
		if err != nil {
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
