package service

import (
	"web_userMessage/user_Message/internal/dao"
	"web_userMessage/user_Message/internal/models"
)

func UpdateUserService(username, age, email, gender string, id int64) (err error) {
	return dao.AlterInformation(username, age, email, gender, id)
}

func GetUserByIdService(phone string) (user *models.User, err error) {
	return dao.GetUser(phone)
}
