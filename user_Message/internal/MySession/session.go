package MySession

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var (
	// Store 定义 Session Store
	Store     = sessions.NewCookieStore([]byte("tL9JQbPqzCjHkK+7DcFjzrEwJ3QoO34="))
	StoreName = "MMM-666"
)

// SetupSession 设置用户 MySession
func SetupSession(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore, storeName, phone string, userId int64) error {
	session, err := store.Get(r, storeName)
	if err != nil {
		log.Println("获取 MySession 失败:", err)
		return err
	}

	// 设置 MySession 值
	session.Values["phone"] = phone
	session.Values["userId"] = userId

	// 配置 MySession 选项
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400, // 一天
		HttpOnly: true,
		Secure:   false, // 开发环境用 false，生产环境用 true
		SameSite: http.SameSiteLaxMode,
	}

	// 保存 MySession
	err = session.Save(r, w)
	if err != nil {
		log.Println("保存 MySession 失败:", err)
		return err
	}

	return nil
}
