package service

import (
	"fmt"
	"web_userMessage/user_Message/internal/dao"
	"web_userMessage/user_Message/internal/models"
	"web_userMessage/user_Message/pkg/validator"
)

// UpdateUserService 更新用户信息服务
func UpdateUserService(username, age, email, gender string, userid int64) error {
	// 参数验证
	if err := validator.ValidateUpdateInput(username, email); err != nil {
		return fmt.Errorf("用户名信息格式有误: %w", err)
	}

	// 调用DAO层更新用户信息
	if err := dao.AlterInformation(username, age, email, gender, userid); err != nil {
		return fmt.Errorf("更新用户信息失败: %w", err)
	}

	return nil
}

func GetUserByIdService(phone string) (user *models.User, err error) {
	return dao.GetUser(phone)
}
