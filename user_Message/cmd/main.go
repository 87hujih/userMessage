package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"web_userMessage/user_Message/config"
	"web_userMessage/user_Message/internal/controllers/admin"
	"web_userMessage/user_Message/internal/controllers/auth"
	"web_userMessage/user_Message/internal/controllers/upload"
	"web_userMessage/user_Message/internal/controllers/user"
	"web_userMessage/user_Message/internal/middleware"
	"web_userMessage/user_Message/pkg/database"
	"web_userMessage/user_Message/pkg/logger"
)

func main() {
	if err := os.MkdirAll("logs", 0755); err != nil {
		fmt.Printf("创建日志目录失败: %v\n", err)
		os.Exit(1)
	}

	// 加载配置
	cfg := config.LoadConfig()
	logger.Infof("配置加载成功，服务器端口: %s", cfg.Server.Port)

	// 代码结束之后关闭数据库连接
	defer func() {
		if err := database.CloseDB(); err != nil {
			logger.Errorf("关闭数据库连接失败: %v", err)
		}
	}()
	server := &http.Server{
		Addr:         "0.0.0.0" + cfg.Server.Port,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}
	//Addr:         ":8090",
	//ReadTimeout:  5 * time.Second,
	//WriteTimeout: 5 * time.Second,

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("user_Message/static"))))
	http.Handle("/user_img/", http.StripPrefix("/user_img/", http.FileServer(http.Dir("user_Message/user_img"))))
	// 公开路由
	http.HandleFunc("/index", user.Index)
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/seekPsd", user.SeekPsd)
	http.HandleFunc("/logout", auth.Logout)

	// 需认证路由
	middleware.RegisterAuthRoute("/homePage", user.HomePage)
	middleware.RegisterAuthRoute("/personalCenter", user.PerCenter)
	middleware.RegisterAuthRoute("/uploadAvatar", upload.Avatar)
	middleware.RegisterAuthRoute("/getUserById", admin.GetUserById)
	middleware.RegisterAuthRoute("/modifyInformation", admin.UpdateUser)
	middleware.RegisterAuthRoute("/deleterUser", admin.DeleterUser)

	logger.Infof("服务器启动在: %s", server.Addr)
	logger.Infof("本地访问: http://localhost%s", cfg.Server.Port)
	logger.Infof("局域网访问: http://192.168.1.100%s", cfg.Server.Port)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panicln(err)
	}

}
