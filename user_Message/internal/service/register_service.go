package service

import "web_userMessage/user_Message/internal/dao"

func RegisterService(name, password, phone string) (err error) {
	return dao.RegisterUser(name, password, phone)
}
