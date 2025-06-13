package logger

import (
	"log"
	"os"
)

var (
	// InfoLogger 信息日志记录器
	InfoLogger *log.Logger
	// ErrorLogger 错误日志记录器
	ErrorLogger *log.Logger
	// DebugLogger 调试日志记录器
	DebugLogger *log.Logger
)

func init() {
	// 创建或打开日志文件
	infoFile, err := os.OpenFile("logs/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 如果无法创建文件，使用标准输出
		InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		InfoLogger = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	errorFile, err := os.OpenFile("logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 如果无法创建文件，使用标准错误输出
		ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		ErrorLogger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	debugFile, err := os.OpenFile("logs/debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 如果无法创建文件，使用标准输出
		DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		DebugLogger = log.New(debugFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

// Info 记录信息日志
func Info(v ...interface{}) {
	InfoLogger.Println(v...)
}

// Infof 格式化记录信息日志
func Infof(format string, v ...interface{}) {
	InfoLogger.Printf(format, v...)
}

// Error 记录错误日志
func Error(v ...interface{}) {
	ErrorLogger.Println(v...)
}

// Errorf 格式化记录错误日志
func Errorf(format string, v ...interface{}) {
	ErrorLogger.Printf(format, v...)
}

// Debug 记录调试日志
func Debug(v ...interface{}) {
	DebugLogger.Println(v...)
}

// Debugf 格式化记录调试日志
func Debugf(format string, v ...interface{}) {
	DebugLogger.Printf(format, v...)
}

// Fatal 记录致命错误并退出程序
func Fatal(v ...interface{}) {
	ErrorLogger.Fatal(v...)
}

// Fatalf 格式化记录致命错误并退出程序
func Fatalf(format string, v ...interface{}) {
	ErrorLogger.Fatalf(format, v...)
}
