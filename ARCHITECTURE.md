# AI Agent 项目架构图

## 系统架构

```
┌─────────────────────────────────────────────────────────────┐
│                        Client Applications                    │
│  (Web Browser / Mobile App / CLI / Postman / etc.)          │
└────────────────────────┬────────────────────────────────────┘
                         │ HTTP/HTTPS Requests
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                     AI Agent Server                          │
│                   (Go + Gin Framework)                       │
│                                                              │
│  ┌──────────────────────────────────────────────────────┐  │
│  │                  HTTP Router Layer                    │  │
│  │  GET  /health              → Health Check Handler    │  │
│  │  POST /api/v1/simple-chat  → Simple Chat Handler     │  │
│  │  POST /api/v1/chat         → Advanced Chat Handler   │  │
│  └────────────────┬─────────────────────────────────────┘  │
│                   │                                          │
│  ┌────────────────▼─────────────────────────────────────┐  │
│  │              Handler Layer                            │  │
│  │  • Request Validation                                │  │
│  │  • Error Handling                                    │  │
│  │  • Response Formatting                               │  │
│  └────────────────┬─────────────────────────────────────┘  │
│                   │                                          │
│  ┌────────────────▼─────────────────────────────────────┐  │
│  │            Agent Service Layer                        │  │
│  │  • Chat(messages []ChatMessage)                      │  │
│  │  • SimpleChat(userMessage string)                    │  │
│  │  • Business Logic                                    │  │
│  └────────────────┬─────────────────────────────────────┘  │
│                   │                                          │
│  ┌────────────────▼─────────────────────────────────────┐  │
│  │           External API Integration                    │  │
│  │  • OpenAI API (GPT-3.5, GPT-4, etc.)                │  │
│  │  •或其他兼容的LLM API                                 │  │
│  └──────────────────────────────────────────────────────┘  │
│                                                              │
│  ┌──────────────────────────────────────────────────────┐  │
│  │            Cross-Cutting Concerns                     │  │
│  │  • Configuration Management (YAML)                   │  │
│  │  • Logging System (Zap)                              │  │
│  │  • Error Handling                                    │  │
│  └──────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                  External Services                          │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │  OpenAI API  │  │  Claude API  │  │  Other LLM   │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
└─────────────────────────────────────────────────────────────┘
```

## 代码组织结构

```
ai-agent/
│
├── cmd/
│   └── server/
│       └── main.go                    # 应用程序入口点
│           ├── 加载配置
│           ├── 初始化日志
│           ├── 创建Agent服务
│           ├── 设置路由
│           └── 启动HTTP服务器
│
├── internal/                           # 私有代码（不可被外部导入）
│   │
│   ├── agent/                         # Agent核心业务逻辑
│   │   ├── agent.go                   # Agent服务实现
│   │   │   ├── NewAgentService()      # 构造函数
│   │   │   ├── Chat()                 # 多消息聊天
│   │   │   └── SimpleChat()           # 简单聊天
│   │   └── agent_test.go              # 单元测试
│   │
│   ├── config/                        # 配置管理
│   │   ├── config.go                  # 配置结构和加载
│   │   │   ├── Config                 # 主配置结构
│   │   │   ├── ServerConfig           # 服务器配置
│   │   │   ├── AgentConfig            # Agent配置
│   │   │   ├── LogConfig              # 日志配置
│   │   │   └── LoadConfig()           # 加载YAML配置
│   │   └── config_test.go             # 单元测试
│   │
│   └── handler/                       # HTTP请求处理
│       └── handler.go                 # 路由和处理器
│           ├── SetupRouter()          # 设置路由
│           ├── handleChat()           # 聊天处理器
│           └── handleSimpleChat()     # 简单聊天处理器
│
├── pkg/                                # 公共库代码
│   └── logger/
│       └── logger.go                  # 日志系统
│           ├── InitLogger()           # 初始化日志
│           ├── Info()                 # 信息日志
│           ├── Error()                # 错误日志
│           ├── Debug()                # 调试日志
│           └── Warn()                 # 警告日志
│
├── configs/
│   └── config.yaml                    # YAML配置文件
│
├── build/                              # 编译输出目录
│   └── server.exe
│
└── 文档和脚本
    ├── README.md
    ├── QUICKSTART.md
    ├── API_EXAMPLES.md
    ├── PROJECT_SUMMARY.md
    ├── CHECKLIST.md
    ├── WELCOME.md
    ├── Makefile
    ├── start.bat
    └── start.ps1
```

## 数据流

### 简单聊天流程

```
Client Request
      │
      ▼
┌─────────────┐
│ HTTP POST   │  POST /api/v1/simple-chat
│  Request    │  Body: {"message": "你好"}
└──────┬──────┘
       │
       ▼
┌─────────────────┐
│ Request Parser  │  验证JSON格式
│  & Validator    │  检查必填字段
└──────┬──────────┘
       │
       ▼
┌─────────────────┐
│   Handler       │  创建消息数组
│  (handler.go)   │  添加system prompt
└──────┬──────────┘
       │
       ▼
┌─────────────────┐
│  Agent Service  │  构建API请求
│  (agent.go)     │  调用LLM API
└──────┬──────────┘
       │
       ▼
┌─────────────────┐
│  External LLM   │  OpenAI/Claude API
│     API         │  返回AI响应
└──────┬──────────┘
       │
       ▼
┌─────────────────┐
│ Response Builder│  格式化响应
│  (handler.go)   │  {"success": true, "data": "..."}
└──────┬──────────┘
       │
       ▼
┌─────────────┐
│ HTTP 200 OK │  JSON Response
│  Response   │
└─────────────┘
```

## 配置层次

```
config.yaml
    │
    ├── server:                          # 服务器配置
    │   ├── host: "0.0.0.0"             # 监听地址
    │   ├── port: 8080                  # 监听端口
    │   ├── read_timeout: 30            # 读取超时(秒)
    │   └── write_timeout: 30           # 写入超时(秒)
    │
    ├── agent:                           # Agent配置
    │   ├── model: "gpt-3.5-turbo"      # LLM模型
    │   ├── api_key: "sk-..."           # API密钥
    │   ├── base_url: "https://..."     # API基础URL
    │   ├── max_tokens: 2000            # 最大token数
    │   ├── temperature: 0.7            # 温度参数
    │   └── system_prompt: "You are..." # 系统提示
    │
    └── log:                             # 日志配置
        ├── level: "info"               # 日志级别
        └── file: "logs/agent.log"      # 日志文件
```

## API端点详情

### 1. 健康检查
```
GET /health

Response:
{
  "status": "ok",
  "time": "2024-01-01T12:00:00Z"
}
```

### 2. 简单聊天
```
POST /api/v1/simple-chat
Content-Type: application/json

Request:
{
  "message": "用户的问题"
}

Response:
{
  "success": true,
  "data": "AI的回答"
}
```

### 3. 高级聊天
```
POST /api/v1/chat
Content-Type: application/json

Request:
{
  "message": "用户的问题"
}

Response:
{
  "success": true,
  "data": "AI的回答"
}

Note: 可扩展为支持多轮对话历史
```

## 扩展点

```
当前架构支持以下扩展：

1. 新的Agent能力
   └─> 在 internal/agent/agent.go 中添加新方法

2. 新的API端点
   └─> 在 internal/handler/handler.go 中添加路由和处理器

3. 新的配置项
   └─> 在 internal/config/config.go 中添加配置结构

4. 中间件支持
   └─> 在 SetupRouter() 中添加中间件
       - 认证中间件
       - 限流中间件
       - CORS中间件
       - 日志中间件

5. 数据库集成
   └─> 添加 database 包
       - 存储对话历史
       - 用户管理
       - 会话管理

6. 缓存层
   └─> 集成Redis
       - 缓存常见响应
       - 会话存储

7. 消息队列
   └─> 集成RabbitMQ/Kafka
       - 异步任务处理
       - 事件驱动架构

8. 监控告警
   └─> 集成Prometheus + Grafana
       - 性能指标
       - 健康监控
```

## 技术栈详情

```
Layer           Technology          Purpose
─────────────────────────────────────────────────
Language        Go 1.21+            编程语言
Web Framework   Gin v1.9.1          HTTP框架
Logging         Zap v1.26.0         结构化日志
Config          yaml.v3 v3.0.1      配置解析
HTTP Client     net/http            标准库HTTP客户端
Testing         testing             标准库测试框架
Build           go build            Go编译器
Package Mgmt    go modules          依赖管理
```

## 部署架构

### 开发环境
```
┌──────────────┐
│ Developer PC │
│  - Go SDK    │
│  - Source    │
│  - Config    │
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ Local Server │
│  Port: 8080  │
└──────────────┘
```

### 生产环境（建议）
```
                    ┌─────────────┐
                    │   Clients   │
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │  Load       │
                    │  Balancer   │
                    └──┬─────┬────┘
                       │     │
              ┌────────▼─┐ ┌─▼────────┐
              │ Server 1 │ │ Server 2 │
              │ :8080    │ │ :8080    │
              └────┬─────┘ └────┬─────┘
                   │            │
              ┌────▼────────────▼────┐
              │   Redis Cache        │
              └──────────┬───────────┘
                         │
              ┌──────────▼───────────┐
              │   Database           │
              │   (PostgreSQL)       │
              └──────────────────────┘
```

---

**架构特点**:
- ✅ 清晰的分层架构
- ✅ 模块化设计
- ✅ 易于扩展
- ✅ 松耦合
- ✅ 可测试性强
