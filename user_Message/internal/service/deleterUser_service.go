package service

import (
	"errors"
	"web_userMessage/user_Message/internal/dao"
)

func DeleterUserService(phone string, phoneNew string) (err error) {
	if phone == phoneNew {
		return errors.New("不能删除自己")
	}
	return dao.DeleteUserByPhone(phone)
}
