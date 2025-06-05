package service

import "web_userMessage/user_Message/internal/dao"

func SeeKPsdService(phone, password string) (err error) {
	return dao.ChangePsd(phone, password)
}
