package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"web_userMessage/user_Message/internal/models"
	"web_userMessage/user_Message/pkg/database"
	"web_userMessage/user_Message/pkg/logger"
	"web_userMessage/user_Message/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
)

// getDB 获取数据库连接
func getDB() *sql.DB {
	return database.GetDB()
}

var db = getDB()

/*注册相关操作*/

// RegisterUser 注册用户
func RegisterUser(username, password, phone string) error {

	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		logger.Errorf("开始事务失败: %v", err)
		return fmt.Errorf("开始事务失败: %w", err)
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				logger.Errorf("回滚事务失败: %v", rollbackErr)
			}
		}
	}()

	// 检查手机号是否已存在
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM users WHERE phone_number = ?", phone).Scan(&count)
	if err != nil {
		logger.Errorf("检查手机号失败: %v", err)
		return fmt.Errorf("检查手机号失败: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("手机号已存在")
	}

	// 插入用户数据
	result, err := tx.Exec("INSERT INTO users (name, password, phone_number) VALUES (?, ?, ?)",
		username, password, phone)
	userID, err := result.LastInsertId()
	_, err = tx.Exec("INSERT INTO information (user_id) VALUES (?)", userID)
	if err != nil {
		logger.Errorf("插入用户数据失败: %v", err)
		return fmt.Errorf("插入用户数据失败: %w", err)
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		logger.Errorf("提交事务失败: %v", err)
		return fmt.Errorf("提交事务失败: %w", err)
	}

	logger.Infof("用户注册成功: %s", phone)
	return nil
}

// CheckNumber 检查手机号是否重复
func CheckNumber(phone string) bool {
	if phone == "" {
		return false
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE phone_number = ?", phone).Scan(&count)
	if err != nil {
		logger.Errorf("检查手机号重复时数据库查询错误: %v", err)
		return false
	}
	return count > 0
}

//--------------------------------------------------------------
/*登录相关操作*/

// LoginUser 用户登录验证
func LoginUser(phone, password string) error {
	if phone == "" || password == "" {
		return fmt.Errorf("手机号和密码不能为空")
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE phone_number = ? AND password = ?", phone, password).Scan(&count)
	if err != nil {
		logger.Errorf("登录验证数据库查询失败: %v", err)
		return fmt.Errorf("登录验证数据库查询失败: %w", err)
	}

	if count == 0 {
		logger.Infof("用户登录失败，用户名或密码错误: %s", phone)
		return utils.ERROR_USER_INFORMATION
	}

	logger.Infof("用户登录成功: %s", phone)
	return nil
}

//--------------------------------------------------------------
/*修改密码相关操作*/

// ChangePsd 修改密码
func ChangePsd(phone, password string) error {
	if !CheckNumber(phone) {
		return utils.ERROR_USER_NOTEXISTS
	}

	_, err := db.Exec("UPDATE users SET password = ? WHERE phone_number = ?", password, phone)
	if err != nil {
		return fmt.Errorf("修改密码失败: %w", err)
	}

	return nil
}

//--------------------------------------------------------------
/*用户信息*/

// GetUser 获取用户信息
func GetUser(phone string) (*models.User, error) {

	u := &models.User{}
	query := `SELECT u.user_id, u.name, u.is_admin, i.age, i.email, i.gender, i.avatar_url 
			  FROM information i 
			  JOIN users u ON i.user_id = u.user_id 
			  WHERE u.phone_number = ?`

	err := db.QueryRow(query, phone).Scan(
		&u.UserId, &u.UserName, &u.IsAdmin,
		&u.Age, &u.Email, &u.Gender, &u.AvatarURL)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.ERROR_USER_NOTEXISTS
		}
		return nil, fmt.Errorf("获取用户信息失败: %w", err)
	}

	return u, nil
}

// GetAllUser 获取指定页的用户列表（分页查询）
func GetAllUser(page, limit int) ([]models.User, error) {
	offset := (page - 1) * limit
	query := `
        SELECT u.user_id, u.name, u.phone_number, u.is_admin, i.age, i.gender, i.avatar_url 
        FROM information i 
        JOIN users u ON i.user_id = u.user_id 
        LIMIT ? OFFSET ?`

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.UserId,
			&user.UserName,
			&user.UserPhone,
			&user.IsAdmin,
			&user.Age,
			&user.Gender,
			&user.AvatarURL); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserCount 获取所有用户数量
func GetUserCount() (count int, err error) {
	err = db.QueryRow("SELECT COUNT(*) AS user_count FROM users").Scan(&count)
	if err != nil {
		log.Println("查询用户数量失败", err)
		return 0, err
	}
	return count, nil
}

// AlterInformation 更新用户信息（使用事务确保数据一致性）
func AlterInformation(username, age, email, gender string, userId int64) error {
	// 验证年龄格式
	var ageInt int
	var err error
	if age != "" {
		ageInt, err = strconv.Atoi(age)
		if err != nil || ageInt < 0 || ageInt > 150 {
			return fmt.Errorf("年龄格式不正确")
		}
	}

	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开始事务失败: %w", err)
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("回滚事务失败: %v", rollbackErr)
			}
		}
	}()

	// 更新用户详细信息
	_, err = tx.Exec("UPDATE information SET age=?, email=?, gender=? WHERE user_id=?", ageInt, email, gender, userId)
	if err != nil {
		return fmt.Errorf("更新用户详细信息失败: %w", err)
	}

	// 更新用户基本信息
	_, err = tx.Exec("UPDATE users SET name=? WHERE user_id=?", username, userId)
	if err != nil {
		return fmt.Errorf("更新用户基本信息失败: %w", err)
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	return nil
}

// UpdateUserAvatar 更新用户头像
func UpdateUserAvatar(userId int64, filename string) error {

	_, err := db.Exec("UPDATE information SET avatar_url=? WHERE user_id=?", filename, userId)
	if err != nil {
		return fmt.Errorf("更新用户头像失败: %w", err)
	}

	return nil
}

// DeleteUserByPhone 删除用户
func DeleteUserByPhone(phone string) error {

	if !CheckNumber(phone) {
		return utils.ERROR_USER_NOTEXISTS
	}

	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开始事务失败: %w", err)
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("回滚事务失败: %v", rollbackErr)
			}
		}
	}()

	// 获取用户ID
	var userId int64
	err = tx.QueryRow("SELECT user_id FROM users WHERE phone_number=?", phone).Scan(&userId)
	if err != nil {
		return fmt.Errorf("获取用户ID失败: %w", err)
	}

	// 删除用户详细信息
	_, err = tx.Exec("DELETE FROM information WHERE user_id=?", userId)
	if err != nil {
		return fmt.Errorf("删除用户详细信息失败: %w", err)
	}

	// 删除用户基本信息
	_, err = tx.Exec("DELETE FROM users WHERE phone_number=?", phone)
	if err != nil {
		return fmt.Errorf("删除用户基本信息失败: %w", err)
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	return nil
}
