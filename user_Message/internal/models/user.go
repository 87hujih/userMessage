package models

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"
	"web_userMessage/user_Message/pkg/utils"
)

type User struct {
	UserId       sql.NullInt64
	UserPhone    sql.NullString
	UserName     sql.NullString
	IsAdmin      sql.NullInt64
	Age          sql.NullInt64
	Gender       sql.NullString
	RegisterData sql.NullString
	Email        sql.NullString
	AvatarURL    sql.NullString
}

var db *sql.DB

func init() {
	err := InitDB() // 调用包中的函数
	db = DB
	if err != nil {
		log.Println("数据库连接失败:", err)
		return
	}
}

/*注册相关操作*/

// RegisterUser  注册
func RegisterUser(username string, password string, phone string) (err error) {
	isPhoneExit := CheckNumber(phone)
	if isPhoneExit {
		return utils.ERROR_USER_EXISTS
	} else {
		userResult, err := db.Exec("insert into users (name,password,phone_number) values (?,?,?)", username, password, phone)
		if err != nil {
			log.Println(err)
			return err
		}
		userID, err := userResult.LastInsertId()
		if err != nil {
			log.Println(err)
			return err

		}
		now := time.Now().Format("2006-01-02")
		_, err = db.Exec("insert into information (user_id,register_date) values (?,?)", userID, now)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

// CheckNumber 检查手机号是否重复
func CheckNumber(phone string) bool {
	var count int
	err := db.QueryRow("select count(*) from users where phone_number=?", phone).Scan(&count)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("数据库查询错误 err=", err)
		return false
	}
	return count > 0
}

//--------------------------------------------------------------
/*登录相关操作*/

func LoginUser(phone string, password string) (err error) {
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE phone_number = ? AND password = ?", phone, password).Scan(&count)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("数据库查询错误 err=", err)
		return err
	}
	if count <= 0 {
		return utils.ERROR_USER_INFORMATION
	}
	return
}

//--------------------------------------------------------------
/*修改密码相关操作*/

// ChangePsd 修改密码
func ChangePsd(phone string, password string) (err error) {
	isPhoneExit := CheckNumber(phone)
	if isPhoneExit {
		_, err := db.Exec("update users set password=? where phone_number=?", password, phone)
		if err != nil {
			log.Println("在数据库修改密码出现问题", err)
		}
		return err
	} else {
		return utils.ERROR_USER_NOTEXISTS
	}
}

//--------------------------------------------------------------
/*用户信息*/

// GetUser 获取到当前用户信息
func GetUser(phone string) (u *User, err error) {
	u = &User{}
	row := db.QueryRow("select u.user_id,u.name,u.is_admin,i.age,i.email,i.gender,i.avatar_url from information i join users u on i.user_id=u.user_id where u.phone_number=?", phone)
	err = row.Scan(&u.UserId, &u.UserName, &u.IsAdmin, &u.Age, &u.Email, &u.Gender, &u.AvatarURL)
	if err != nil {
		log.Println(err)
		return &User{}, err
	}
	return u, nil
}

// GetAllUser 获取指定页的用户列表（分页查询）
func GetAllUser(page, limit int) ([]User, error) {
	offset := (page - 1) * limit
	query := `
        SELECT u.user_id, u.name, u.phone_number, u.is_admin, i.age, i.gender, i.register_date, i.avatar_url 
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

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.UserId,
			&user.UserName,
			&user.UserPhone,
			&user.IsAdmin,
			&user.Age,
			&user.Gender,
			&user.RegisterData,
			&user.AvatarURL); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetAllUser 获取所有用户信息
//func GetAllUser() ([]User, error) {
//	rows, err := db.Query("select u.user_id,u.name,u.phone_number,u.is_admin,i.age,i.gender,i.register_date,i.avatar_url from information i join users u on i.user_id=u.user_id")
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}
//	var users []User
//	for rows.Next() {
//		var user User
//		if err := rows.Scan(&user.UserId, &user.UserName, &user.UserPhone, &user.IsAdmin, &user.Age, &user.Gender, &user.RegisterData, &user.AvatarURL); err != nil {
//			fmt.Println(err)
//			return nil, err
//		}
//		users = append(users, user)
//	}
//	return users, nil
//}

// GetUserCount 获取所有用户数量
func GetUserCount() (count int, err error) {
	err = db.QueryRow("SELECT COUNT(*) AS user_count FROM users").Scan(&count)
	if err != nil {
		log.Println("查询用户数量失败", err)
		return 0, err
	}
	return count, nil
}

// AlterInformation 保存个人信息
func AlterInformation(username string, age string, email string, gender string, userId int64) (err error) {
	num, _ := strconv.Atoi(age) // 将字符串转换为 int 类型
	_, err = db.Exec("UPDATE information SET age=?, email=? ,gender=? WHERE user_id=?", num, email, gender, userId)
	_, err = db.Exec("UPDATE users SET  name=? WHERE user_id=?", username, userId)
	return err
}

// UpdateUserAvatar 修改个人头像
func UpdateUserAvatar(userId int64, filename string) (err error) {
	_, err = db.Exec("UPDATE information SET avatar_url=? WHERE user_id=?", filename, userId)
	return err
}

// DeleteUserByPhone 删除用户
func DeleteUserByPhone(phone string) (err error) {
	_, err = db.Exec("DELETE from users where phone_number=?", phone)
	return err
}
