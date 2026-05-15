# Test AI Agent API

Write-Host "Testing AI Agent API" -ForegroundColor Cyan
Write-Host ""

# Test 1: Health Check
Write-Host "[Test 1] Health Check..." -ForegroundColor Green
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/health" -Method GET
    Write-Host "Success!" -ForegroundColor Green
    Write-Host ($response | ConvertTo-Json)
} catch {
    Write-Host "Failed: $_" -ForegroundColor Red
    exit 1
}

Write-Host ""

# Test 2: Simple Chat
Write-Host "[Test 2] Simple Chat..." -ForegroundColor Green
try {
    $body = '{"message": "Hello, please introduce yourself"}'
    Write-Host "Sending request..." -ForegroundColor Yellow
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/simple-chat" -Method POST -ContentType "application/json" -Body $body
    
    if ($response.success) {
        Write-Host "Success!" -ForegroundColor Green
        Write-Host "AI Response:" -ForegroundColor Cyan
        Write-Host $response.data
    } else {
        Write-Host "Failed: $($response.error)" -ForegroundColor Red
    }
} catch {
    Write-Host "Request failed: $_" -ForegroundColor Red
}

Write-Host ""
Write-Host "All tests completed!" -ForegroundColor Cyan
