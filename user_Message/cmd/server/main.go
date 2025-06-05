package main

import (
	"log"
	"net/http"
	"time"
	"web_userMessage/user_Message/internal/controllers/admin"
	"web_userMessage/user_Message/internal/controllers/auth"
	"web_userMessage/user_Message/internal/controllers/upload"
	"web_userMessage/user_Message/internal/controllers/user"
	"web_userMessage/user_Message/internal/middleware"
)

func main() {
	server := http.Server{
		Addr:         ":8090",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("user_Message/static"))))
	http.Handle("/user_img/", http.StripPrefix("/user_img/", http.FileServer(http.Dir("user_Message/user_img"))))
	//公开路由
	http.HandleFunc("/index", user.Index)
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/seekPsd", user.SeekPsd)
	http.HandleFunc("/logout", auth.Logout)
	//需认证路由
	http.Handle("/homePage", middleware.Auth(middleware.Store, middleware.StoreName, "/login")(http.HandlerFunc(user.HomePage)))
	http.Handle("/personalCenter", middleware.Auth(middleware.Store, middleware.StoreName, "/login")(http.HandlerFunc(user.PerCenter)))
	http.Handle("/uploadAvatar", middleware.Auth(middleware.Store, middleware.StoreName, "/login")(http.HandlerFunc(upload.UploadAvatar)))
	http.Handle("/getUserById", middleware.Auth(middleware.Store, middleware.StoreName, "/login")(http.HandlerFunc(admin.GetUserById)))
	http.Handle("/modifyInformation", middleware.Auth(middleware.Store, middleware.StoreName, "/login")(http.HandlerFunc(admin.UpdateUser)))
	http.Handle("/deleterUser", middleware.Auth(middleware.Store, middleware.StoreName, "/login")(http.HandlerFunc(admin.DeleterUser)))
	//http.HandleFunc("/homePage", user.HomePage)
	//http.HandleFunc("/personalCenter", user.PerCenter)
	//http.HandleFunc("/uploadAvatar", upload.UploadAvatar)
	//http.HandleFunc("/getUserById", admin.GetUserById)
	//http.HandleFunc("/modifyInformation", admin.UpdateUser)
	//http.HandleFunc("/deleterUser", admin.DeleterUser)

	err := server.ListenAndServe()
	if err != nil {
		log.Panicln(err)
		return
	}
}
