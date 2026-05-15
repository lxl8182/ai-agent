# 测试游戏攻略搜索功能

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Testing Game Guide Search" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

$baseUrl = "http://localhost:8080"

# 测试1: 原神攻略
Write-Host "[Test 1] Genshin Impact - Raiden Shogun Guide..." -ForegroundColor Yellow
$body1 = @{
    message = "原神 雷电将军培养攻略"
} | ConvertTo-Json -Depth 10

try {
    $response1 = Invoke-RestMethod -Uri "$baseUrl/api/v1/smart-chat" `
        -Method Post `
        -ContentType "application/json; charset=utf-8" `
        -Body $body1
    
    Write-Host "✓ Success!" -ForegroundColor Green
    Write-Host "Response:" -ForegroundColor Cyan
    Write-Host $response1.data -ForegroundColor White
} catch {
    Write-Host "✗ Failed: $_" -ForegroundColor Red
}

Write-Host ""
Write-Host "Press Enter to continue..." -ForegroundColor Gray
Read-Host

# 测试2: 黑神话悟空
Write-Host "[Test 2] Black Myth Wukong - Boss Strategy..." -ForegroundColor Yellow
$body2 = @{
    message = "黑神话悟空 虎先锋打法"
} | ConvertTo-Json -Depth 10

try {
    $response2 = Invoke-RestMethod -Uri "$baseUrl/api/v1/smart-chat" `
        -Method Post `
        -ContentType "application/json; charset=utf-8" `
        -Body $body2
    
    Write-Host "✓ Success!" -ForegroundColor Green
    Write-Host "Response:" -ForegroundColor Cyan
    Write-Host $response2.data -ForegroundColor White
} catch {
    Write-Host "✗ Failed: $_" -ForegroundColor Red
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Tests Completed!" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
