package middleware

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

// Auth 认证中间件，检查用户是否登录
func Auth(store *sessions.CookieStore, storeName string, loginURL string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 获取 session
			sess, err := store.Get(r, storeName)
			if err != nil {
				log.Println("获取 session 失败:", err)
				http.Error(w, "内部服务器错误", http.StatusInternalServerError)
				return
			}

			// 检查用户是否登录
			phone, ok := sess.Values["phone"]
			if !ok || phone == nil {
				// 未登录，重定向到登录页面
				http.Redirect(w, r, loginURL, http.StatusFound)
				return
			}
			// 已登录，继续处理请求
			//log.Printf("用户已登录: %v", phone)
			next.ServeHTTP(w, r)
		})
	}
}

// RegisterAuthRoute 用户状态认证
func RegisterAuthRoute(pattern string, handler http.HandlerFunc) {
	http.Handle(pattern, Auth(Store, StoreName, "/login")(handler))
}
