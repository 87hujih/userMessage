package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入但不直接使用，用于注册驱动
	"time"
)

var DB *sql.DB

// InitDB 数据库连接池
func InitDB() (err error) {
	dbUser := "root"
	dbPass := "654321"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "ums"
	// 构建连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// 连接数据库
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error()) // 无法连接数据库
		return err
	}

	DB.SetMaxOpenConns(10)                 // 最大打开连接数
	DB.SetMaxIdleConns(5)                  // 最大空闲连接数
	DB.SetConnMaxLifetime(5 * time.Minute) // 连接最大存活时间
	// 测试连接是否成功
	err = DB.Ping()
	if err != nil {
		panic(err.Error()) // 无法 ping 通数据库
		return err
	}
	return
	//fmt.Println("成功连接到 MySQL 数据库")
}
