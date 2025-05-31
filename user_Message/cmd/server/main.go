package main

import (
	"log"
	"net/http"
	"time"
	"web_userMessage/user_Message/internal/handlers"
)

func main() {
	server := http.Server{
		Addr:         ":8090",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("user_Message/static"))))
	http.Handle("/user_img/", http.StripPrefix("/user_img/", http.FileServer(http.Dir("user_Message/user_img"))))
	http.HandleFunc("/index", handlers.Index)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/seekPsd", handlers.SeekPsd)
	http.HandleFunc("/homePage", handlers.HomePage)
	http.HandleFunc("/personalCenter", handlers.PerCenter)
	http.HandleFunc("/uploadAvatar", handlers.UploadAvatar)
	http.HandleFunc("/getUserById", handlers.GetUserById)
	http.HandleFunc("/modifyInformation", handlers.UpdateUser)
	http.HandleFunc("/deleterUser", handlers.DeleterUser)
	http.HandleFunc("/logout", handlers.Logout)
	err := server.ListenAndServe()
	if err != nil {
		log.Panicln(err)
		return
	}
}
