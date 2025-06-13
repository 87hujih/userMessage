package service

import (
	"web_userMessage/user_Message/internal/dao"
	"web_userMessage/user_Message/pkg/validator"
)

// RegisterService 用户注册服务
func RegisterService(name, password, phone string) error {
	// 参数验证
	if err := validator.ValidateUserInput(name, password, phone); err != nil {
		return err
	}

	// 调用DAO层注册用户
	if err := dao.RegisterUser(name, password, phone); err != nil {
		return err
	}

	return nil
}
