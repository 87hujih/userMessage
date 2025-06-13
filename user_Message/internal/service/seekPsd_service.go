package service

import (
	"fmt"
	"web_userMessage/user_Message/internal/dao"
	"web_userMessage/user_Message/pkg/validator"
)

// SeeKPsdService 修改密码服务
func SeeKPsdService(phone, password string) error {
	// 参数验证
	if phone == "" {
		return fmt.Errorf("手机号不能为空")
	}
	if password == "" {
		return fmt.Errorf("新密码不能为空")
	}

	// 验证手机号格式
	if !validator.IsValidPhone(phone) {
		return fmt.Errorf("手机号格式不正确")
	}

	// 验证密码强度
	if len(password) < 6 {
		return fmt.Errorf("密码长度不能少于6位")
	}

	// 调用DAO层修改密码
	if err := dao.ChangePsd(phone, password); err != nil {
		return err
	}

	return nil
}
