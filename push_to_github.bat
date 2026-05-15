@echo off
chcp 65001 >nul
echo ========================================
echo   Pushing to GitHub
echo ========================================
echo.
echo Repository: https://github.com/lxl8182/ai-agent.git
echo.
echo Please make sure you have created the repository on GitHub first!
echo.
pause

cd /d %~dp0

echo Pushing to GitHub...
git push -u origin main

echo.
echo ========================================
if %ERRORLEVEL% EQU 0 (
    echo Success! Check your repository at:
    echo https://github.com/lxl8182/ai-agent
) else (
    echo Failed! Please check the error message above.
)
echo ========================================
pause
