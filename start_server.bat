@echo off
chcp 65001 >nul
echo ========================================
echo          Go Web 服务器启动脚本
echo ========================================
echo.

echo [1/3] 检查防火墙规则...
netsh advfirewall firewall show rule name="Go Web Server" >nul 2>&1
if %errorlevel% neq 0 (
    echo 添加防火墙规则...
    netsh advfirewall firewall add rule name="Go Web Server" dir=in action=allow protocol=TCP localport=8090
    echo 防火墙规则添加成功！
) else (
    echo 防火墙规则已存在
)

echo.
echo [2/3] 获取本机IP地址...
echo 本机可用的IP地址：
for /f "tokens=2 delims=:" %%i in ('ipconfig ^| findstr /i "IPv4"') do (
    for /f "tokens=1" %%j in ("%%i") do (
        echo   - %%j
    )
)

echo.
echo [3/3] 启动服务器...
echo 服务器将在以下地址启动：
echo   - 本地访问: http://localhost:8090
echo   - 局域网访问: http://你的IP地址:8090
echo.
echo 按任意键启动服务器...
pause >nul

cd user_Message
go run cmd/main.go

echo.
echo 服务器已停止
pause