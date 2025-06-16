# 外网访问配置指南

## 概述

本指南将帮助您配置项目，使其可以通过URL在别人的电脑上访问。

## 当前问题分析

### 🔍 发现的问题
1. **硬编码监听地址**: `cmd/main.go` 中服务器地址硬编码为 `:8090`
2. **未使用配置文件**: 没有使用 `cfg.Server.Port` 配置
3. **仅本地访问**: 当前配置只能本机访问

## 解决方案

### 方案一：局域网访问（推荐用于开发/测试）

#### 1. 修改服务器监听地址
需要将服务器从只监听本地改为监听所有网络接口。

**修改 `cmd/main.go`**：
```go
// 将这行
server := &http.Server{
    Addr:         ":8090",  // 只监听本地
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 5 * time.Second,
}

// 改为
server := &http.Server{
    Addr:         "0.0.0.0" + cfg.Server.Port,  // 监听所有网络接口
    ReadTimeout:  cfg.Server.ReadTimeout,
    WriteTimeout: cfg.Server.WriteTimeout,
}
```

#### 2. 配置防火墙
**Windows 防火墙设置**：
```powershell
# 允许端口 8090 通过防火墙
netsh advfirewall firewall add rule name="Go Web Server" dir=in action=allow protocol=TCP localport=8090
```

#### 3. 获取本机IP地址
```powershell
# 查看本机IP地址
ipconfig
```

#### 4. 访问方式
- **局域网内访问**: `http://你的IP地址:8090`
- **例如**: `http://192.168.1.100:8090`

### 方案二：公网访问（生产环境）

#### 1. 云服务器部署
**推荐云服务商**：
- 阿里云 ECS
- 腾讯云 CVM
- 华为云 ECS
- AWS EC2

#### 2. 域名配置
```bash
# 购买域名并配置DNS解析
# 将域名指向服务器公网IP
```

#### 3. HTTPS配置
```go
// 生产环境建议使用HTTPS
server := &http.Server{
    Addr:         ":443",
    ReadTimeout:  cfg.Server.ReadTimeout,
    WriteTimeout: cfg.Server.WriteTimeout,
}

// 使用SSL证书
log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
```

#### 4. 反向代理（推荐）
**使用 Nginx 配置**：
```nginx
server {
    listen 80;
    server_name yourdomain.com;
    
    location / {
        proxy_pass http://localhost:8090;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

### 方案三：内网穿透（临时测试）

#### 1. 使用 ngrok
```bash
# 下载并安装 ngrok
# 启动内网穿透
ngrok http 8090
```

#### 2. 使用花生壳
- 下载花生壳客户端
- 配置内网穿透
- 获得公网访问地址

## 代码修改建议

### 1. 修改 `cmd/main.go`
```go
package main

import (
    "errors"
    "fmt"
    "log"
    "net/http"
    "os"
    "web_userMessage/user_Message/config"
    // ... 其他导入
)

func main() {
    // ... 现有代码 ...
    
    // 修改服务器配置
    server := &http.Server{
        Addr:         "0.0.0.0" + cfg.Server.Port,  // 监听所有接口
        ReadTimeout:  cfg.Server.ReadTimeout,
        WriteTimeout: cfg.Server.WriteTimeout,
        IdleTimeout:  cfg.Server.IdleTimeout,
    }
    
    // ... 路由配置 ...
    
    logger.Infof("服务器启动在: %s", server.Addr)
    logger.Infof("访问地址: http://localhost%s", cfg.Server.Port)
    
    if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
        log.Panicln(err)
    }
}
```

### 2. 环境变量配置
**修改 `.env` 文件**：
```env
# 服务器配置
SERVER_PORT=:8090
SERVER_HOST=0.0.0.0  # 新增：监听所有接口
SERVER_READ_TIMEOUT=5s
SERVER_WRITE_TIMEOUT=5s
SERVER_IDLE_TIMEOUT=5s
```

### 3. 配置结构体更新
**修改 `config/config.go`**：
```go
type ServerConfig struct {
    Host         string        `json:"host"`          // 新增
    Port         string        `json:"port"`
    ReadTimeout  time.Duration `json:"read_timeout"`
    WriteTimeout time.Duration `json:"write_timeout"`
    IdleTimeout  time.Duration `json:"idle_timeout"`
    StaticDir    string        `json:"static_dir"`
    UploadDir    string        `json:"upload_dir"`
}

// 在 LoadConfig 函数中添加
Server: ServerConfig{
    Host:         getEnv("SERVER_HOST", "0.0.0.0"),  // 新增
    Port:         getEnv("SERVER_PORT", ":8090"),
    // ... 其他配置
}
```

## 安全注意事项

### ⚠️ 重要提醒
1. **生产环境必须使用HTTPS**
2. **配置防火墙规则**
3. **定期更新依赖包**
4. **使用强密码和密钥**
5. **启用访问日志监控**

### 🔒 安全配置
```env
# 生产环境配置
SESSION_SECURE=true     # 启用安全Cookie
SESSION_HTTP_ONLY=true  # 防止XSS
```

## 测试步骤

### 1. 本地测试
```bash
# 启动服务器
go run cmd/main.go

# 测试本地访问
curl http://localhost:8090
```

### 2. 局域网测试
```bash
# 在其他设备上测试
curl http://你的IP:8090
```

### 3. 公网测试
```bash
# 使用公网IP或域名测试
curl http://你的域名或公网IP:8090
```

## 常见问题

### Q: 无法访问怎么办？
A: 检查以下项目：
1. 防火墙是否开放端口
2. 服务器是否正确监听 0.0.0.0
3. 网络连接是否正常
4. 端口是否被占用

### Q: 如何查看服务器日志？
A: 查看 `logs/` 目录下的日志文件

### Q: 性能优化建议？
A: 
1. 使用反向代理（Nginx）
2. 启用Gzip压缩
3. 配置静态文件缓存
4. 使用CDN加速

---

**配置完成后，您的项目就可以通过URL在别人的电脑上访问了！**