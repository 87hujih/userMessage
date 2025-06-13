package service

import (
	"fmt"
	"web_userMessage/user_Message/internal/dao"
	"web_userMessage/user_Message/pkg/validator"
)

// UploadAvatarService 上传头像服务
func UploadAvatarService(userId int64, filename string) error {
	// 参数验证
	if userId <= 0 {
		return fmt.Errorf("无效的用户ID")
	}
	if filename == "" {
		return fmt.Errorf("文件名不能为空")
	}

	// 验证文件扩展名
	if !validator.IsValidImageFile(filename) {
		return fmt.Errorf("不支持的图片格式，仅支持 jpg, jpeg, png, gif")
	}

	// 调用DAO层更新用户头像
	if err := dao.UpdateUserAvatar(userId, filename); err != nil {
		return fmt.Errorf("更新用户头像失败: %w", err)
	}

	return nil
}
