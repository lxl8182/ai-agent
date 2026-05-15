# 环境变量配置说明

为了保护API密钥等敏感信息，建议使用环境变量而不是硬编码在配置文件中。

## 设置环境变量

### Windows (PowerShell)
```powershell
$env:AGENT_API_KEY="sk-373e1f6b941f4ef4958a9ccd72b75039"
$env:AGENT_MODEL="qwen-turbo"
$env:AGENT_BASE_URL="https://dashscope.aliyuncs.com/compatible-mode/v1"
$env:SERVER_PORT="8080"
```

### Windows (CMD)
```batch
set AGENT_API_KEY=sk-373e1f6b941f4ef4958a9ccd72b75039
set AGENT_MODEL=qwen-turbo
set AGENT_BASE_URL=https://dashscope.aliyuncs.com/compatible-mode/v1
set SERVER_PORT=8080
```

### Linux/Mac
```bash
export AGENT_API_KEY="sk-373e1f6b941f4ef4958a9ccd72b75039"
export AGENT_MODEL="qwen-turbo"
export AGENT_BASE_URL="https://dashscope.aliyuncs.com/compatible-mode/v1"
export SERVER_PORT="8080"
```

## 使用 .env 文件（推荐）

1. 创建 `.env` 文件（已在 .gitignore 中排除）
2. 添加以下内容：
```env
AGENT_API_KEY=sk-373e1f6b941f4ef4958a9ccd72b75039
AGENT_MODEL=qwen-turbo
AGENT_BASE_URL=https://dashscope.aliyuncs.com/compatible-mode/v1
SERVER_PORT=8080
```

3. 运行前加载环境变量：
```bash
# Linux/Mac
source .env

# Windows PowerShell
Get-Content .env | ForEach-Object {
    if ($_ -match '^([^#][^=]+)=(.*)$') {
        [Environment]::SetEnvironmentVariable($matches[1], $matches[2])
    }
}
```

## 可用的环境变量

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| AGENT_API_KEY | API密钥 | 必填 |
| AGENT_MODEL | 模型名称 | qwen-turbo |
| AGENT_BASE_URL | API基础URL | https://dashscope.aliyuncs.com/compatible-mode/v1 |
| AGENT_MAX_TOKENS | 最大token数 | 2000 |
| AGENT_TEMPERATURE | 温度参数 | 0.7 |
| AGENT_SYSTEM_PROMPT | 系统提示词 | 你是一个有用的AI助手 |
| SERVER_HOST | 服务器地址 | 0.0.0.0 |
| SERVER_PORT | 服务器端口 | 8080 |
| LOG_LEVEL | 日志级别 | info |
| LOG_FILE | 日志文件路径 | logs/agent.log |
