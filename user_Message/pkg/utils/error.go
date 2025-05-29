package utils

import "errors"

var (
	ERROR_USER_NOTEXISTS   = errors.New("该用户不存在..")
	ERROR_USER_EXISTS      = errors.New("该手机号已经被注册..")
	ERROR_USER_INFORMATION = errors.New("账号或密码错误..")
)
