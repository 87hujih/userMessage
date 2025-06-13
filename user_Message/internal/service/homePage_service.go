package service

import (
	"fmt"
	"web_userMessage/user_Message/internal/dao"
	"web_userMessage/user_Message/internal/models"
)

func HomePageService(phone string, page, limit int) (error, []models.User, *models.User, int) {
	var allUser []models.User
	user := &models.User{}
	if phone == "" {
		return fmt.Errorf("手机号和密码不能为空"), allUser, user, 0
	}
	if page == 0 || limit == 0 {
		return fmt.Errorf("页数和查询参数不能为0"), allUser, user, 0
	}
	user, err := dao.GetUser(phone)
	if err != nil {
		return fmt.Errorf("获取用户不存在"), allUser, user, 0
	}
	allUser, err = dao.GetAllUser(page, limit)
	if err != nil {
		return fmt.Errorf("获取全部用户失败"), allUser, user, 0
	}
	total, err := dao.GetUserCount()
	if err != nil {
		return fmt.Errorf("获取全部用户数量失败"), allUser, user, 0
	}
	return nil, allUser, user, total
}
