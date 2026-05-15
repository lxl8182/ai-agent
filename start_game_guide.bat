@echo off
chcp 65001 >nul
echo ========================================
echo   Game Guide Agent - Quick Start
echo ========================================
echo.
echo This script will help you set up and run the Game Guide Agent.
echo.

REM Check if config file exists
if not exist "configs\config.yaml" (
    echo [INFO] Config file not found. Creating from template...
    copy "configs\config.yaml.example" "configs\config.yaml"
    echo [INFO] Please edit configs\config.yaml and add your API key.
    echo.
)

REM Prompt for Bing API Key
echo [SETUP] Do you want to configure Bing Search API?
echo   - Get your API key from: https://portal.azure.com
echo   - Or skip and use without search feature
echo.
set /p setup_bing="Enter Bing API Key (or press Enter to skip): "

if "%setup_bing%"=="" (
    echo [INFO] Skipping Bing API configuration.
    echo [WARN] Search feature will be disabled.
    echo.
) else (
    echo [INFO] Setting BING_SEARCH_API_KEY environment variable...
    setx BING_SEARCH_API_KEY "%setup_bing%"
    set BING_SEARCH_API_KEY=%setup_bing%
    echo [SUCCESS] Bing API Key configured!
    echo.
)

echo ========================================
echo   Starting Server...
echo ========================================
echo.
echo Server will start on http://localhost:8080
echo.
echo Available endpoints:
echo   - GET  /health                    (Health check)
echo   - POST /api/v1/simple-chat        (Simple chat)
echo   - POST /api/v1/smart-chat         (Smart chat with search) ⭐
echo.
echo Example usage:
echo   curl -X POST http://localhost:8080/api/v1/smart-chat ^
echo     -H "Content-Type: application/json" ^
echo     -d "{\"message\": \"原神 雷电将军怎么培养\"}"
echo.
echo Press Ctrl+C to stop the server.
echo.

go run cmd/server/main.go
