package service

import "web_userMessage/user_Message/internal/dao"

func UploadAvatarService(userId int64, filename string) (err error) {
	return dao.UpdateUserAvatar(userId, filename)
}
