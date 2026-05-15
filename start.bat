@echo off
echo ========================================
echo   AI Agent Server - Quick Start
echo ========================================
echo.

REM 检查配置文件
if not exist "configs\config.yaml" (
    echo [ERROR] Configuration file not found: configs\config.yaml
    echo Please create the configuration file first.
    pause
    exit /b 1
)

echo [INFO] Checking Go installation...
go version >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Go is not installed or not in PATH
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

echo [INFO] Go installation found
echo.

echo [INFO] Building server...
go build -o build\server.exe cmd\server\main.go
if errorlevel 1 (
    echo [ERROR] Build failed
    pause
    exit /b 1
)

echo [INFO] Build successful
echo.

echo [INFO] Starting AI Agent Server...
echo [INFO] Press Ctrl+C to stop the server
echo.

build\server.exe
