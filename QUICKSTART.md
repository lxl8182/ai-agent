# 快速开始指南

## 🚀 5分钟快速启动AI Agent

### 步骤1: 配置API密钥

编辑 `configs/config.yaml` 文件，填入你的API密钥：

```yaml
agent:
  api_key: "sk-your-actual-api-key-here"  # ← 替换这里
  model: "gpt-3.5-turbo"
  base_url: "https://api.openai.com/v1"
```

**获取API密钥：**
- OpenAI: https://platform.openai.com/api-keys
- 或其他兼容的LLM提供商

### 步骤2: 运行服务

在项目根目录执行：

```bash
# 方法1: 直接运行
go run cmd/server/main.go

# 方法2: 使用Makefile
make run
```

你会看到类似输出：
```
2024-01-01T12:00:00.000+0800    INFO    Starting AI Agent Server...
2024-01-01T12:00:00.000+0800    INFO    Server starting on 0.0.0.0:8080
```

### 步骤3: 测试API

打开新终端，测试健康检查：

```bash
curl http://localhost:8080/health
```

测试聊天功能：

```bash
curl -X POST http://localhost:8080/api/v1/simple-chat \
  -H "Content-Type: application/json" \
  -d '{"message": "你好！"}'
```

### 步骤4: 开始开发

#### 添加自定义功能

1. **扩展Agent能力** - 编辑 `internal/agent/agent.go`：

```go
func (a *AgentService) AnalyzeText(text string) (string, error) {
    prompt := fmt.Sprintf("请分析以下文本的情感倾向：%s", text)
    return a.SimpleChat(prompt)
}
```

2. **添加新API端点** - 编辑 `internal/handler/handler.go`：

```go
api.POST("/analyze", handleAnalyze(agentService))
```

3. **添加中间件** - 在 `SetupRouter` 中添加：

```go
router.Use(AuthMiddleware())
```

## 🔧 常用命令

```bash
# 构建项目
make build

# 运行测试
make test

# 格式化代码
make fmt

# 清理构建文件
make clean

# 查看帮助
make help
```

## 📝 项目结构速览

```
ai-agent/
├── cmd/server/main.go          # 程序入口
├── configs/config.yaml         # 配置文件 ← 首先编辑这个
├── internal/
│   ├── agent/agent.go         # Agent核心逻辑
│   ├── config/config.go       # 配置管理
│   └── handler/handler.go     # HTTP路由和处理器
├── pkg/logger/logger.go        # 日志模块
└── README.md                   # 完整文档
```

## 🎯 下一步

- 📖 查看 [README.md](README.md) 了解完整功能
- 🔌 查看 [API_EXAMPLES.md](API_EXAMPLES.md) 学习API使用
- 🧪 运行 `make test` 确保一切正常
- 🚀 开始构建你的AI应用！

## ❓ 常见问题

**Q: 如何更换AI模型？**
A: 修改 `configs/config.yaml` 中的 `model` 字段

**Q: 支持哪些模型？**
A: 任何OpenAI兼容的API，包括GPT-3.5、GPT-4、Claude等

**Q: 如何查看日志？**
A: 日志默认输出到控制台和 `logs/agent.log`

**Q: 如何修改端口？**
A: 修改 `configs/config.yaml` 中的 `server.port`

## 💡 提示

- 首次运行时，确保网络连接正常
- API密钥请妥善保管，不要提交到版本控制
- 建议使用 `.env` 文件管理敏感信息（需要自行实现）
- 生产环境请配置合适的日志级别和超时时间

---

准备好了吗？开始构建你的AI Agent应用吧！🎉
