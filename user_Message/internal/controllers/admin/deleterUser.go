package admin

import (
	"log"
	"net/http"
	"web_userMessage/user_Message/internal/controllers/user"
	md "web_userMessage/user_Message/internal/models"
	cm "web_userMessage/user_Message/pkg/utils"
)

// DeleterUser 删除用户
func DeleterUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	phone := r.FormValue("phone")
	session, err := user.Store.Get(r, user.StoreName)
	phoneNew, _ := session.Values["phone"].(string)
	if phone == phoneNew {
		cm.SendMessage(w, 500, "不能删除自己")
		log.Println(err)
		return
	}
	err = md.DeleteUserByPhone(phone)
	if err != nil {
		cm.SendMessage(w, 500, "删除用户失败")
		return
	}
	cm.SendMessage(w, 200, "删除成功")
}
