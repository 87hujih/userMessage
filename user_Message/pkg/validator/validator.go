package validator

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

// IsValidPhone 验证手机号格式
func IsValidPhone(phone string) bool {
	// 简单的手机号验证正则表达式（11位数字，以1开头）
	pattern := `^1[3-9]\d{9}$`
	matched, _ := regexp.MatchString(pattern, phone)
	return matched
}

// IsValidEmail 验证邮箱格式
func IsValidEmail(email string) bool {
	// 简单的邮箱验证正则表达式
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// IsValidImageFile 验证是否为有效的图片文件
func IsValidImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	validExts := []string{".jpg", ".jpeg", ".png", ".gif"}

	for _, validExt := range validExts {
		if ext == validExt {
			return true
		}
	}
	return false
}

// ValidateUserInput 验证用户输入的基本信息
func ValidateUserInput(name, password, phone string) error {
	if name == "" {
		return fmt.Errorf("用户名不能为空")
	}
	if password == "" {
		return fmt.Errorf("密码不能为空")
	}
	if phone == "" {
		return fmt.Errorf("手机号不能为空")
	}

	// 验证手机号格式
	if !IsValidPhone(phone) {
		fmt.Println(phone)
		return fmt.Errorf("手机号格式不正确")
	}

	// 验证密码强度
	if len(password) < 6 {
		return fmt.Errorf("密码长度不能少于6位")
	}

	// 验证用户名长度
	if len(name) < 2 || len(name) > 20 {
		return fmt.Errorf("用户名长度应在2-20个字符之间")
	}

	return nil
}

// ValidateUpdateInput 验证更新用户信息的输入
func ValidateUpdateInput(username, email string) error {
	// 验证用户名
	if username != "" && (len(username) < 2 || len(username) > 20) {
		return fmt.Errorf("用户名长度应在2-20个字符之间")
	}
	// 验证邮箱格式
	if email != "" && !IsValidEmail(email) {
		return fmt.Errorf("邮箱格式不正确")
	}
	return nil
}
