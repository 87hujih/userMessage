package main

import (
	"log"
	"net/http"
	"time"
	"web_userMessage/user_Message/internal/controllers"
)

func main() {
	server := http.Server{
		Addr:         ":8090",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("user_Message/static"))))
	http.Handle("/user_img/", http.StripPrefix("/user_img/", http.FileServer(http.Dir("user_Message/user_img"))))
	http.HandleFunc("/index", controllers.Index)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/seekPsd", controllers.SeekPsd)
	http.HandleFunc("/homePage", controllers.HomePage)
	http.HandleFunc("/personalCenter", controllers.PerCenter)
	http.HandleFunc("/uploadAvatar", controllers.UploadAvatar)
	http.HandleFunc("/getUserById", controllers.GetUserById)
	http.HandleFunc("/modifyInformation", controllers.UpdateUser)
	http.HandleFunc("/deleterUser", controllers.DeleterUser)
	http.HandleFunc("/logout", controllers.Logout)
	err := server.ListenAndServe()
	if err != nil {
		log.Panicln(err)
		return
	}
}
