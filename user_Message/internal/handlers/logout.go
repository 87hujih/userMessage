package handlers

import "net/http"

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
