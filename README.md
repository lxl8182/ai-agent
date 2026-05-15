# AI Agent Project

一个基于Go语言的AI Agent服务项目，提供RESTful API接口与大语言模型进行交互。

## 项目结构

```
ai-agent/
├── cmd/server/          # 主程序入口
├── configs/             # 配置文件
├── internal/            # 内部包
│   ├── agent/          # Agent核心逻辑
│   ├── config/         # 配置管理
│   └── handler/        # HTTP处理器
├── pkg/                # 公共包
│   └── logger/         # 日志模块
├── go.mod              # Go模块文件
└── README.md           # 项目说明
```

## 功能特性

- 🤖 支持OpenAI兼容的API接口
- 🌐 RESTful API设计
- ⚙️ YAML配置文件管理
- 📝 结构化日志记录
- 🔧 可扩展的Agent架构
- 🎮 **新增：游戏攻略搜索智能体**（自动搜索最新攻略）
- 🔍 **新增：网络搜索工具集成**（Bing/Google）
- 💡 **新增：智能意图识别**（自动判断是否需要搜索）

## 快速开始

### 1. 环境要求

- Go 1.21+
- OpenAI API Key（或其他兼容的LLM API）

### 2. 配置API密钥

**重要：** 为了保护你的API密钥，有两种方式：

#### 方式1：使用配置文件（本地开发）
```bash
# 复制配置模板
cp configs/config.yaml.example configs/config.yaml

# 编辑配置文件，填入你的API密钥
# 注意：configs/config.yaml 已在 .gitignore 中，不会被提交到Git
```

#### 方式2：使用环境变量（推荐，更安全）
```bash
# Windows PowerShell
$env:AGENT_API_KEY="sk-your-api-key-here"

# Linux/Mac
export AGENT_API_KEY="sk-your-api-key-here"
```

详见 [ENV_SETUP.md](ENV_SETUP.md) 和 [GITHUB_UPLOAD_GUIDE.md](GITHUB_UPLOAD_GUIDE.md)

### 3. 安装依赖

```bash
go mod tidy
```

### 4. 运行服务

```bash
go run cmd/server/main.go
```

服务将在 `http://localhost:8080` 启动

**注意：** 如果使用环境变量配置，确保在运行前已设置 `AGENT_API_KEY` 等变量。

## API接口

### 健康检查
```
GET /health
```

### 简单聊天
```
POST /api/v1/simple-chat
Content-Type: application/json

{
  "message": "你好，请介绍一下你自己"
}
```

### 高级聊天（支持多轮对话）
```
POST /api/v1/chat
Content-Type: application/json

{
  "message": "你的消息内容"
}
```

### ⭐ 智能对话（推荐 - 支持网络搜索）
```
POST /api/v1/smart-chat
Content-Type: application/json

{
  "message": "原神 雷电将军怎么培养"
}
```

**特点**:
- ✅ 自动识别游戏相关问题
- ✅ 调用 Bing/Google 搜索最新攻略
- ✅ 整合搜索结果生成详细回答
- ✅ 非游戏问题直接回答，不调用搜索

**示例**:
```bash
curl -X POST http://localhost:8080/api/v1/smart-chat \
  -H "Content-Type: application/json" \
  -d '{"message": "黑神话悟空 虎先锋怎么打"}'
```

详见 [GAME_GUIDE_AGENT.md](GAME_GUIDE_AGENT.md)

## 配置说明

### 服务器配置
- `host`: 监听地址
- `port`: 监听端口
- `read_timeout`: 读取超时时间（秒）
- `write_timeout`: 写入超时时间（秒）

### Agent配置
- `model`: LLM模型名称
- `api_key`: API密钥
- `base_url`: API基础URL
- `max_tokens`: 最大token数
- `temperature`: 温度参数（0-1）
- `system_prompt`: 系统提示词

### 日志配置
- `level`: 日志级别（debug/info/warn/error）
- `file`: 日志文件路径

## 扩展开发

### 添加新的Agent功能

在 `internal/agent/agent.go` 中添加新方法：

```go
func (a *AgentService) YourNewFunction(params string) (string, error) {
    // 实现你的逻辑
}
```

### 添加新的API端点

在 `internal/handler/handler.go` 中添加新的处理函数和路由。

## 部署

### Docker部署（可选）

创建 `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY configs ./configs
EXPOSE 8080
CMD ["./server"]
```

## 许可证

MIT License
