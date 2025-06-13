package admin

import (
	"log"
	"net/http"
	"web_userMessage/user_Message/internal/middleware"
	"web_userMessage/user_Message/internal/service"
	"web_userMessage/user_Message/pkg/utils"
)

// DeleterUser 删除用户
func DeleterUser(w http.ResponseWriter, r *http.Request) {
	session, err := middleware.Store.Get(r, middleware.StoreName)
	phoneNew, ok := session.Values["phone"].(string)

	if !ok || err != nil {
		http.Error(w, "用户未登录", http.StatusUnauthorized)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	phone := r.FormValue("phone")
	//处理删除业务
	err = service.DeleterUserService(phone, phoneNew)

	if err != nil {
		utils.SendMessage(w, 500, "删除用户失败")
		return
	}
	utils.SendMessage(w, 200, "删除成功")
}
