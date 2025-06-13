package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

// Config 应用配置结构
type Config struct {
	Database DatabaseConfig `json:"database"`
	Server   ServerConfig   `json:"server"`
	Session  SessionConfig  `json:"session"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host            string        `json:"host"`
	Port            string        `json:"port"`
	User            string        `json:"user"`
	Password        string        `json:"password"`
	DBName          string        `json:"db_name"`
	MaxOpenConns    int           `json:"max_open_coins"`
	MaxIdleConns    int           `json:"max_idle_coins"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Host         string        `json:"host"`
	Port         string        `json:"port"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	IdleTimeout  time.Duration `json:"idle_timeout"`
	StaticDir    string        `json:"static_dir"`
	UploadDir    string        `json:"upload_dir"`
}

// SessionConfig Session配置
type SessionConfig struct {
	SecretKey string `json:"secret_key"`
	Name      string `json:"name"`
	MaxAge    int    `json:"max_age"`
	Secure    bool   `json:"secure"`
	HttpOnly  bool   `json:"http_only"`
}

// LoadConfig 加载配置
func LoadConfig() *Config {
	//读取项目跟目录下的.env文件把里面的配置加载为环境变量
	err := godotenv.Load()
	if err != nil {
		fmt.Println("未找到 .env 文件，继续使用默认或系统环境变量")
	}
	return &Config{
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "3306"),
			User:            getEnv("DB_USER", "root"),
			Password:        getEnv("DB_PASSWORD", "654321"),
			DBName:          getEnv("DB_NAME", "ums"),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
		},
		Server: ServerConfig{
			Host:         getEnv("SERVER_HOST", "0.0.0.0"),
			Port:         getEnv("SERVER_PORT", ":8090"),
			ReadTimeout:  getEnvAsDuration("SERVER_READ_TIMEOUT", 5*time.Second),
			WriteTimeout: getEnvAsDuration("SERVER_WRITE_TIMEOUT", 5*time.Second),
			IdleTimeout:  getEnvAsDuration("SERVER_IDLE_TIMEOUT", 5*time.Second),
			StaticDir:    getEnv("STATIC_DIR", "user_Message/static"),
			UploadDir:    getEnv("UPLOAD_DIR", "user_Message/user_img"),
		},
		Session: SessionConfig{
			SecretKey: getEnvRequired("SESSION_SECRET"),
			Name:      getEnv("SESSION_NAME", "secure-session"),
			MaxAge:    getEnvAsInt("SESSION_MAX_AGE", 7200),
			Secure:    getEnvAsBool("SESSION_SECURE", false),
			HttpOnly:  getEnvAsBool("SESSION_HTTP_ONLY", true),
		},
	}
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}

// 优先读取环境变量，如果没有设置，就使用默认值

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

// getEnvRequired 获取必需的环境变量
func getEnvRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("必需的环境变量 %s 未设置，请在 .env 文件中配置", key)
	}
	return value
}
