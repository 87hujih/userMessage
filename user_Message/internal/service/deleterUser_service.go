package service

import (
	"fmt"
	"web_userMessage/user_Message/internal/dao"
)

// DeleterUserService 删除用户服务
func DeleterUserService(phone string, currentUserPhone string) error {
	// 参数验证
	if phone == "" {
		return fmt.Errorf("要删除的用户手机号不能为空")
	}
	if currentUserPhone == "" {
		return fmt.Errorf("当前用户手机号不能为空")
	}

	// 不能删除自己
	if phone == currentUserPhone {
		return fmt.Errorf("不能删除自己")
	}

	// 调用DAO层删除用户
	if err := dao.DeleteUserByPhone(phone); err != nil {
		return fmt.Errorf("删除用户失败: %w", err)
	}

	return nil
}
