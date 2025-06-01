package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"web_userMessage/user_Message/internal/controllers/user"
	md "web_userMessage/user_Message/internal/models"
	cm "web_userMessage/user_Message/pkg/utils"
)

// UploadAvatar 头像上传处理
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "仅支持 POST 请求", http.StatusMethodNotAllowed)
		return
	}
	// 设置最大上传大小为 5MB
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		http.Error(w, "解析表单失败", http.StatusBadRequest)
		return
	}
	// 获取上传的文件
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "无法获取文件", http.StatusBadRequest)
		return
	}
	defer func() {
		_ = file.Close()
	}()
	// 检查文件类型
	if !strings.HasPrefix(handler.Header.Get("Content-Type"), "image/") {
		http.Error(w, "仅支持图片格式", http.StatusBadRequest)
		return
	}
	session, err := user.Store.Get(r, user.StoreName)
	if err != nil {
		http.Error(w, "会话无效", http.StatusInternalServerError)
		return
	}
	userId, ok := session.Values["userId"].(int64)
	if !ok {
		http.Error(w, "用户未登录", http.StatusUnauthorized)
		return
	}
	// 生成唯一文件名
	filename := generateUniqueFilename(handler.Filename)
	// 保存文件
	filePath := filepath.Join(user.AvatarDir, filename)
	if err := saveUploadedFile(file, filePath); err != nil {
		http.Error(w, "文件保存失败", http.StatusInternalServerError)
		return
	}

	// 更新数据库
	if err := md.UpdateUserAvatar(userId, filename); err != nil {
		http.Error(w, "更新数据库失败", http.StatusInternalServerError)
		return
	}

	cm.SendMessage(w, 200, "头像上传成功")
}

// generateUniqueFilename 生成唯一的文件名
func generateUniqueFilename(originalName string) string {
	ext := filepath.Ext(originalName)
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}

// saveUploadedFile 将上传的文件写入磁盘
func saveUploadedFile(src multipart.File, dst string) error {
	outFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		_ = outFile.Close()
	}()

	_, err = io.Copy(outFile, src)
	return err
}
