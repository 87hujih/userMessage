package service

import (
	"fmt"
	"web_userMessage/user_Message/internal/dao"
	"web_userMessage/user_Message/pkg/utils"
	"web_userMessage/user_Message/pkg/validator"
)

// LoginService 用户登录服务
func LoginService(phone, password string) (error, int64) {
	if phone == "" || password == "" {
		return fmt.Errorf("手机号和密码不能为空"), 0
	}

	// 验证手机号格式
	if !validator.IsValidPhone(phone) {
		return fmt.Errorf("手机号格式不正确"), 0
	}

	// 验证用户登录
	if err := dao.LoginUser(phone, password); err != nil {
		return utils.ERROR_USER_INFORMATION, 0
	}

	// 获取用户信息
	user, err := dao.GetUser(phone)

	if err != nil {
		return fmt.Errorf("获取用户信息失败: %w", err), 0
	}

	if !user.UserId.Valid {
		return fmt.Errorf("用户ID无效"), 0
	}

	return nil, user.UserId.Int64
}
