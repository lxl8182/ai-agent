@echo off
chcp 65001 >nul
echo ========================================
echo   Testing AI Agent with Alibaba Cloud
echo ========================================
echo.

echo [Test 1] Health Check...
curl http://localhost:8080/health
echo.
echo.

echo [Test 2] Chat with AI (English)...
curl -X POST http://localhost:8080/api/v1/simple-chat -H "Content-Type: application/json" -d "{\"message\": \"Hello, please introduce yourself\"}"
echo.
echo.

echo [Test 3] Another question...
curl -X POST http://localhost:8080/api/v1/simple-chat -H "Content-Type: application/json" -d "{\"message\": \"What is AI?\"}"
echo.
echo.

echo ========================================
echo   Tests completed!
echo ========================================
pause
