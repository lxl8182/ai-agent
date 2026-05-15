# 游戏攻略智能体测试脚本
# 使用前请确保服务已启动：go run cmd/server/main.go

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Game Guide Agent Test Script" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

$baseUrl = "http://localhost:8080"

# 测试1: 健康检查
Write-Host "[Test 1] Health Check..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$baseUrl/health" -Method Get
    Write-Host "✓ Health check passed" -ForegroundColor Green
    Write-Host "Response: $($response | ConvertTo-Json)" -ForegroundColor Gray
} catch {
    Write-Host "✗ Health check failed: $_" -ForegroundColor Red
}
Write-Host ""

# 测试2: 原神攻略查询
Write-Host "[Test 2] Genshin Impact Character Guide..." -ForegroundColor Yellow
try {
    $body = @{
        message = "原神 雷电将军怎么培养"
    } | ConvertTo-Json
    
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/smart-chat" `
        -Method Post `
        -ContentType "application/json" `
        -Body $body
    
    Write-Host "✓ Request successful" -ForegroundColor Green
    Write-Host "Response:" -ForegroundColor Cyan
    Write-Host $response.data -ForegroundColor White
} catch {
    Write-Host "✗ Request failed: $_" -ForegroundColor Red
}
Write-Host ""
Write-Host "Press Enter to continue to next test..." -ForegroundColor Gray
Read-Host

# 测试3: Boss打法查询
Write-Host "[Test 3] Boss Strategy Guide..." -ForegroundColor Yellow
try {
    $body = @{
        message = "黑神话悟空 虎先锋怎么打"
    } | ConvertTo-Json
    
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/smart-chat" `
        -Method Post `
        -ContentType "application/json" `
        -Body $body
    
    Write-Host "✓ Request successful" -ForegroundColor Green
    Write-Host "Response:" -ForegroundColor Cyan
    Write-Host $response.data -ForegroundColor White
} catch {
    Write-Host "✗ Request failed: $_" -ForegroundColor Red
}
Write-Host ""
Write-Host "Press Enter to continue to next test..." -ForegroundColor Gray
Read-Host

# 测试4: 非游戏问题（应该不调用搜索）
Write-Host "[Test 4] Non-game Question (should not use search)..." -ForegroundColor Yellow
try {
    $body = @{
        message = "你好，请介绍一下你自己"
    } | ConvertTo-Json
    
    $response = Invoke-RestMethod -Uri "$baseUrl/api/v1/smart-chat" `
        -Method Post `
        -ContentType "application/json" `
        -Body $body
    
    Write-Host "✓ Request successful" -ForegroundColor Green
    Write-Host "Response:" -ForegroundColor Cyan
    Write-Host $response.data -ForegroundColor White
} catch {
    Write-Host "✗ Request failed: $_" -ForegroundColor Red
}
Write-Host ""

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  All tests completed!" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Note: If search is not working, make sure BING_SEARCH_API_KEY is set" -ForegroundColor Yellow
Write-Host ""
