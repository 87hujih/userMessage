package database

import (
	"database/sql"
	"fmt"
	"time"
	"web_userMessage/user_Message/config"
	"web_userMessage/user_Message/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// DB 全局数据库连接池
	DB *sql.DB
)

// init 初始化数据库连接池
func init() {
	cfg := config.LoadConfig()
	// 构建数据源名称
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Errorf("打开数据库连接失败: %w", err)
	}

	// 设置连接池参数
	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime * time.Minute)

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		fmt.Errorf("数据库连接测试失败: %w", err)
	}
	DB = db
	logger.Info("数据库连接池初始化成功")
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if DB != nil {
		if err := DB.Close(); err != nil {
			return fmt.Errorf("关闭数据库连接失败: %w", err)
		}
		logger.Info("数据库连接已关闭")
	}
	return nil
}

// GetDB 获取数据库连接
func GetDB() *sql.DB {
	return DB
}

// Ping 检查数据库连接状态
func Ping() error {
	if DB == nil {
		return fmt.Errorf("数据库连接未初始化")
	}
	return DB.Ping()
}

// Stats 获取数据库连接池统计信息
func Stats() sql.DBStats {
	if DB == nil {
		return sql.DBStats{}
	}
	return DB.Stats()
}
