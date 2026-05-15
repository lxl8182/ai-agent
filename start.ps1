# AI Agent Server - Quick Start Script for PowerShell

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  AI Agent Server - Quick Start" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 检查配置文件
if (-not (Test-Path "configs\config.yaml")) {
    Write-Host "[ERROR] Configuration file not found: configs\config.yaml" -ForegroundColor Red
    Write-Host "Please create the configuration file first." -ForegroundColor Yellow
    Read-Host "Press Enter to exit"
    exit 1
}

# 检查Go安装
Write-Host "[INFO] Checking Go installation..." -ForegroundColor Green
try {
    $goVersion = go version 2>&1
    Write-Host "[INFO] Go installation found: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "[ERROR] Go is not installed or not in PATH" -ForegroundColor Red
    Write-Host "Please install Go from https://golang.org/dl/" -ForegroundColor Yellow
    Read-Host "Press Enter to exit"
    exit 1
}

Write-Host ""

# 构建服务器
Write-Host "[INFO] Building server..." -ForegroundColor Green
go build -o build\server.exe cmd\server\main.go

if ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] Build failed" -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

Write-Host "[INFO] Build successful" -ForegroundColor Green
Write-Host ""

# 启动服务器
Write-Host "[INFO] Starting AI Agent Server..." -ForegroundColor Green
Write-Host "[INFO] Press Ctrl+C to stop the server" -ForegroundColor Yellow
Write-Host ""

& ".\build\server.exe"
