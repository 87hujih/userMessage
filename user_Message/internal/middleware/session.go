package middleware

import (
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"web_userMessage/user_Message/config"
)

var (
	// Store 定义 Session Store
	Store     *sessions.CookieStore
	StoreName string
)

func init() {
	cfg := config.LoadConfig()
	Store = sessions.NewCookieStore([]byte(cfg.Session.SecretKey))
	StoreName = cfg.Session.Name
}

// SetupSession 设置用户Session
func SetupSession(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore, storeName, phone string, userId int64) error {
	if phone == "" {
		return fmt.Errorf("手机号不能为空")
	}
	if userId <= 0 {
		return fmt.Errorf("无效的用户ID")
	}

	session, err := store.Get(r, storeName)
	if err != nil {
		log.Printf("获取Session失败: %v", err)
		return fmt.Errorf("获取Session失败: %w", err)
	}

	// 设置Session值
	session.Values["phone"] = phone
	session.Values["userId"] = userId

	// 从配置文件获取Session选项
	cfg := config.LoadConfig()
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   cfg.Session.MaxAge,
		HttpOnly: cfg.Session.HttpOnly,
		Secure:   cfg.Session.Secure,
		SameSite: http.SameSiteLaxMode,
	}

	// 保存Session
	if err = session.Save(r, w); err != nil {
		log.Printf("保存Session失败: %v", err)
		return fmt.Errorf("保存Session失败: %w", err)
	}

	return nil
}
