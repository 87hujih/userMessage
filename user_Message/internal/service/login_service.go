package service

import (
	"web_userMessage/user_Message/internal/dao"
)

func LoginService(phone, password string) (err error, id int64) {
	err = dao.LoginUser(phone, password)
	if err != nil {
		return err, 0
	}
	u, err := dao.GetUser(phone)
	return err, u.UserId.Int64
}
