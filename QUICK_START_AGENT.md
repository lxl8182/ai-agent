# 🚀 智能体升级 - 快速实施指南

## 第一阶段：会话管理 + 基础工具（1-2周）

### Step 1: 添加依赖

```bash
# 进入项目目录
cd E:\job\code\ai-agent

# 添加 Redis 客户端
go get github.com/redis/go-redis/v9

# 添加 UUID 生成
go get github.com/google/uuid

# 更新依赖
go mod tidy
```

---

### Step 2: 创建会话管理模块

创建文件 `internal/session/session_manager.go`:

```go
package session

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// ChatMessage 对话消息
type ChatMessage struct {
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// Session 会话
type Session struct {
	ID        string        `json:"id"`
	UserID    string        `json:"user_id"`
	Messages  []ChatMessage `json:"messages"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// Manager 会话管理器接口
type Manager interface {
	CreateSession(userID string) (*Session, error)
	GetSession(sessionID string) (*Session, error)
	AddMessage(sessionID string, message ChatMessage) error
	GetHistory(sessionID string, limit int) ([]ChatMessage, error)
	DeleteSession(sessionID string) error
}

// RedisManager Redis 实现的会话管理器
type RedisManager struct {
	client *redis.Client
	ttl    time.Duration
}

// NewRedisManager 创建 Redis 会话管理器
func NewRedisManager(addr string, password string, db int, ttl time.Duration) (*RedisManager, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// 测试连接
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisManager{
		client: client,
		ttl:    ttl,
	}, nil
}

// CreateSession 创建新会话
func (m *RedisManager) CreateSession(userID string) (*Session, error) {
	session := &Session{
		ID:        uuid.New().String(),
		UserID:    userID,
		Messages:  make([]ChatMessage, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	ctx := context.Background()
	key := m.getSessionKey(session.ID)

	data, err := json.Marshal(session)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal session: %w", err)
	}

	if err := m.client.Set(ctx, key, data, m.ttl).Err(); err != nil {
		return nil, fmt.Errorf("failed to save session: %w", err)
	}

	return session, nil
}

// GetSession 获取会话
func (m *RedisManager) GetSession(sessionID string) (*Session, error) {
	ctx := context.Background()
	key := m.getSessionKey(sessionID)

	data, err := m.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("session not found")
		}
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	var session Session
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &session, nil
}

// AddMessage 添加消息到会话
func (m *RedisManager) AddMessage(sessionID string, message ChatMessage) error {
	session, err := m.GetSession(sessionID)
	if err != nil {
		return err
	}

	message.Timestamp = time.Now()
	session.Messages = append(session.Messages, message)
	session.UpdatedAt = time.Now()

	ctx := context.Background()
	key := m.getSessionKey(sessionID)

	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	if err := m.client.Set(ctx, key, data, m.ttl).Err(); err != nil {
		return fmt.Errorf("failed to update session: %w", err)
	}

	return nil
}

// GetHistory 获取会话历史
func (m *RedisManager) GetHistory(sessionID string, limit int) ([]ChatMessage, error) {
	session, err := m.GetSession(sessionID)
	if err != nil {
		return nil, err
	}

	messages := session.Messages
	if len(messages) > limit {
		messages = messages[len(messages)-limit:]
	}

	return messages, nil
}

// DeleteSession 删除会话
func (m *RedisManager) DeleteSession(sessionID string) error {
	ctx := context.Background()
	key := m.getSessionKey(sessionID)

	if err := m.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	return nil
}

func (m *RedisManager) getSessionKey(sessionID string) string {
	return fmt.Sprintf("session:%s", sessionID)
}
```

---

### Step 3: 创建工具系统

创建文件 `internal/tools/tool_registry.go`:

```go
package tools

import (
	"fmt"
	"sync"
)

// Tool 工具接口
type Tool interface {
	Name() string
	Description() string
	Parameters() map[string]interface{}
	Execute(args map[string]interface{}) (string, error)
}

// Registry 工具注册中心
type Registry struct {
	tools map[string]Tool
	mu    sync.RWMutex
}

// NewRegistry 创建工具注册中心
func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]Tool),
	}
}

// Register 注册工具
func (r *Registry) Register(tool Tool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tools[tool.Name()] = tool
}

// Execute 执行工具
func (r *Registry) Execute(name string, args map[string]interface{}) (string, error) {
	r.mu.RLock()
	tool, exists := r.tools[name]
	r.mu.RUnlock()

	if !exists {
		return "", fmt.Errorf("tool %s not found", name)
	}

	return tool.Execute(args)
}

// ListTools 列出所有可用工具
func (r *Registry) ListTools() []map[string]interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []map[string]interface{}
	for _, tool := range r.tools {
		result = append(result, map[string]interface{}{
			"name":        tool.Name(),
			"description": tool.Description(),
			"parameters":  tool.Parameters(),
		})
	}

	return result
}
```

---

### Step 4: 实现基础工具

创建文件 `internal/tools/calculator.go`:

```go
package tools

import (
	"fmt"
	"go/token"
	"go/types"

	"go/constant"
	"go/parser"
)

// CalculatorTool 计算器工具
type CalculatorTool struct{}

func (t *CalculatorTool) Name() string {
	return "calculator"
}

func (t *CalculatorTool) Description() string {
	return "执行数学计算，支持加减乘除、幂运算等"
}

func (t *CalculatorTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"expression": map[string]interface{}{
			"type":        "string",
			"description": "数学表达式，例如: 2 + 2 * 3",
			"required":    true,
		},
	}
}

func (t *CalculatorTool) Execute(args map[string]interface{}) (string, error) {
	expr, ok := args["expression"].(string)
	if !ok {
		return "", fmt.Errorf("expression is required")
	}

	result, err := evaluateExpression(expr)
	if err != nil {
		return "", fmt.Errorf("calculation error: %w", err)
	}

	return fmt.Sprintf("计算结果: %s", result), nil
}

func evaluateExpression(expr string) (string, error) {
	// 简单的表达式求值（生产环境建议使用更安全的方案）
	exp, err := parser.ParseExpr(expr)
	if err != nil {
		return "", err
	}

	tv := types.EvalConstant(nil, exp)
	if tv.Kind() != constant.Int && tv.Kind() != constant.Float {
		return "", fmt.Errorf("unsupported expression type")
	}

	return tv.String(), nil
}
```

创建文件 `internal/tools/datetime.go`:

```go
package tools

import (
	"fmt"
	"time"
)

// DateTimeTool 日期时间工具
type DateTimeTool struct{}

func (t *DateTimeTool) Name() string {
	return "datetime"
}

func (t *DateTimeTool) Description() string {
	return "获取当前日期和时间信息"
}

func (t *DateTimeTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"format": map[string]interface{}{
			"type":        "string",
			"description": "日期格式，默认: 2006-01-02 15:04:05",
			"required":    false,
		},
		"timezone": map[string]interface{}{
			"type":        "string",
			"description": "时区，例如: Asia/Shanghai",
			"required":    false,
		},
	}
}

func (t *DateTimeTool) Execute(args map[string]interface{}) (string, error) {
	format := "2006-01-02 15:04:05"
	if f, ok := args["format"].(string); ok && f != "" {
		format = f
	}

	location := time.Local
	if tz, ok := args["timezone"].(string); ok && tz != "" {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			return "", fmt.Errorf("invalid timezone: %w", err)
		}
		location = loc
	}

	now := time.Now().In(location)
	return fmt.Sprintf("当前时间: %s", now.Format(format)), nil
}
```

---

### Step 5: 增强 Agent 服务

修改 `internal/agent/agent.go`，添加工具调用能力：

```go
package agent

import (
	"ai-agent/internal/config"
	"ai-agent/internal/tools"
	"encoding/json"
	"fmt"
	"strings"
)

type IntelligentAgent struct {
	config       config.AgentConfig
	toolRegistry *tools.Registry
}

func NewIntelligentAgent(cfg config.AgentConfig, registry *tools.Registry) *IntelligentAgent {
	return &IntelligentAgent{
		config:       cfg,
		toolRegistry: registry,
	}
}

// SmartChat 智能对话（支持工具调用）
func (a *IntelligentAgent) SmartChat(message string, history []ChatMessage) (string, error) {
	// 1. 构建提示词，包含可用工具信息
	toolsInfo := a.getToolsInfo()
	
	systemPrompt := fmt.Sprintf(`%s

你有以下工具可用：
%s

如果需要调用工具，请使用以下格式：
TOOL_CALL: {"tool": "工具名", "args": {"参数": "值"}}

否则直接回答问题。`, a.config.SystemPrompt, toolsInfo)

	// 2. 构建消息列表
	messages := []ChatMessage{
		{Role: "system", Content: systemPrompt},
	}
	
	// 添加历史记录
	messages = append(messages, history...)
	messages = append(messages, ChatMessage{Role: "user", Content: message})

	// 3. 调用 LLM
	response, err := a.Chat(messages)
	if err != nil {
		return "", err
	}

	// 4. 检查是否需要调用工具
	if strings.Contains(response, "TOOL_CALL:") {
		toolCall := extractToolCall(response)
		if toolCall != nil {
			// 执行工具
			result, err := a.toolRegistry.Execute(toolCall.Tool, toolCall.Args)
			if err != nil {
				return "", fmt.Errorf("tool execution failed: %w", err)
			}

			// 基于工具结果生成最终回答
			finalPrompt := fmt.Sprintf(
				"工具执行结果: %s\n请基于此结果回答用户的问题: %s",
				result,
				message,
			)
			
			return a.SimpleChat(finalPrompt)
		}
	}

	// 5. 直接返回回答
	return response, nil
}

type ToolCall struct {
	Tool string                 `json:"tool"`
	Args map[string]interface{} `json:"args"`
}

func (a *IntelligentAgent) getToolsInfo() string {
	toolsList := a.toolRegistry.ListTools()
	
	var sb strings.Builder
	for i, tool := range toolsList {
		sb.WriteString(fmt.Sprintf("%d. %s: %s\n", i+1, tool["name"], tool["description"]))
	}
	
	return sb.String()
}

func extractToolCall(response string) *ToolCall {
	idx := strings.Index(response, "TOOL_CALL:")
	if idx == -1 {
		return nil
	}

	jsonStr := strings.TrimSpace(response[idx+len("TOOL_CALL:"):])
	
	var toolCall ToolCall
	if err := json.Unmarshal([]byte(jsonStr), &toolCall); err != nil {
		return nil
	}

	return &toolCall
}
```

---

### Step 6: 更新 Handler

修改 `internal/handler/handler.go`，添加新的 API 端点：

```go
package handler

import (
	"ai-agent/internal/agent"
	"ai-agent/internal/config"
	"ai-agent/internal/session"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	agent        *agent.IntelligentAgent
	sessionMgr   session.Manager
	serverConfig config.ServerConfig
}

func NewHandler(
	agent *agent.IntelligentAgent,
	sessionMgr session.Manager,
	serverConfig config.ServerConfig,
) *Handler {
	return &Handler{
		agent:        agent,
		sessionMgr:   sessionMgr,
		serverConfig: serverConfig,
	}
}

func SetupRoutes(
	router *gin.Engine,
	agent *agent.IntelligentAgent,
	sessionMgr session.Manager,
	serverConfig config.ServerConfig,
) {
	handler := NewHandler(agent, sessionMgr, serverConfig)

	// 健康检查
	router.GET("/health", handler.healthCheck)

	// API v1
	v1 := router.Group("/api/v1")
	{
		v1.POST("/simple-chat", handler.simpleChat)
		v1.POST("/chat", handler.chat)
		v1.POST("/smart-chat", handler.smartChat) // ⭐ 新增
		v1.GET("/tools", handler.listTools)        // ⭐ 新增
		v1.GET("/session/:id", handler.getSession) // ⭐ 新增
	}
}

// smartChat 智能对话
func (h *Handler) smartChat(c *gin.Context) {
	var req struct {
		Message string `json:"message" binding:"required"`
		UserID  string `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request: " + err.Error(),
		})
		return
	}

	if req.UserID == "" {
		req.UserID = "anonymous"
	}

	// 获取或创建会话
	sessionID := c.GetHeader("X-Session-ID")
	var sess *session.Session
	var err error

	if sessionID != "" {
		sess, err = h.sessionMgr.GetSession(sessionID)
		if err != nil {
			// 会话不存在，创建新的
			sess, _ = h.sessionMgr.CreateSession(req.UserID)
			sessionID = sess.ID
		}
	} else {
		sess, _ = h.sessionMgr.CreateSession(req.UserID)
		sessionID = sess.ID
	}

	// 获取历史消息
	history, _ := h.sessionMgr.GetHistory(sessionID, 10)

	// 执行智能对话
	response, err := h.agent.SmartChat(req.Message, history)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// 保存消息到会话
	h.sessionMgr.AddMessage(sessionID, session.ChatMessage{
		Role:    "user",
		Content: req.Message,
	})
	h.sessionMgr.AddMessage(sessionID, session.ChatMessage{
		Role:    "assistant",
		Content: response,
	})

	// 返回响应
	c.Header("X-Session-ID", sessionID)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": map[string]interface{}{
			"response":   response,
			"session_id": sessionID,
		},
	})
}

// listTools 列出可用工具
func (h *Handler) listTools(c *gin.Context) {
	tools := h.agent.ListTools()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    tools,
	})
}

// getSession 获取会话信息
func (h *Handler) getSession(c *gin.Context) {
	sessionID := c.Param("id")
	
	sess, err := h.sessionMgr.GetSession(sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Session not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": map[string]interface{}{
			"id":         sess.ID,
			"user_id":    sess.UserID,
			"created_at": sess.CreatedAt,
			"msg_count":  len(sess.Messages),
		},
	})
}

// ... 其他现有方法保持不变
```

---

### Step 7: 更新主程序

修改 `cmd/server/main.go`:

```go
package main

import (
	"ai-agent/internal/agent"
	"ai-agent/internal/config"
	"ai-agent/internal/handler"
	"ai-agent/internal/session"
	"ai-agent/internal/tools"
	"ai-agent/pkg/logger"
	"fmt"
	"log"
	"time"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志
	if err := logger.InitLogger(cfg.Log.Level, cfg.Log.File); err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Starting AI Agent Server...")

	// 初始化工具注册中心
	toolRegistry := tools.NewRegistry()
	toolRegistry.Register(&tools.CalculatorTool{})
	toolRegistry.Register(&tools.DateTimeTool{})
	// 添加更多工具...

	// 初始化会话管理器
	sessionMgr, err := session.NewRedisManager(
		"localhost:6379", // Redis 地址
		"",               // 密码
		0,                // 数据库
		*time.Hour,       // TTL
	)
	if err != nil {
		logger.Warn(fmt.Sprintf("Redis not available, using in-memory session: %v", err))
		// TODO: 实现内存版会话管理器作为 fallback
	}

	// 初始化智能 Agent
	agentService := agent.NewIntelligentAgent(cfg.Agent, toolRegistry)

	// 设置路由
	router := handler.SetupRouter(agentService, sessionMgr, cfg.Server)

	// 启动服务器
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	logger.Info(fmt.Sprintf("Server starting on %s", addr))
	
	if err := router.Run(addr); err != nil {
		logger.Error(fmt.Sprintf("Failed to start server: %v", err))
	}
}
```

---

### Step 8: 更新配置文件

修改 `configs/config.yaml.example`:

```yaml
server:
  host: "0.0.0.0"
  port: 8080
  read_timeout: 30
  write_timeout: 30

agent:
  model: "qwen-turbo"
  api_key: "${AGENT_API_KEY}"
  base_url: "https://dashscope.aliyuncs.com/api/v1"
  max_tokens: 2000
  temperature: 0.7
  system_prompt: |
    You are a helpful AI assistant with access to various tools.
    When you need to use a tool, respond with:
    TOOL_CALL: {"tool": "tool_name", "args": {"arg1": "value1"}}

log:
  level: "info"
  file: "logs/agent.log"

# ⭐ 新增：会话配置
session:
  store: "redis"
  redis_addr: "localhost:6379"
  ttl: 3600
```

---

### Step 9: 安装和测试

#### 1. 安装 Redis

**Windows**:
```powershell
# 使用 Chocolatey
choco install redis-64

# 或使用 Docker
docker run -d -p 6379:6379 redis:latest
```

**Linux/Mac**:
```bash
# Ubuntu/Debian
sudo apt-get install redis-server

# Mac
brew install redis
```

#### 2. 启动 Redis

```bash
# 直接启动
redis-server

# 或后台运行
redis-server --daemonize yes
```

#### 3. 编译和运行

```bash
# 编译
go build -o build/server cmd/server/main.go

# 运行
./build/server

# 或使用 Make
make run
```

#### 4. 测试 API

```powershell
# 测试健康检查
curl http://localhost:8080/health

# 测试智能对话
curl -X POST http://localhost:8080/api/v1/smart-chat `
  -H "Content-Type: application/json" `
  -d '{
    "message": "计算 25 * 48 的结果",
    "user_id": "test-user"
  }'

# 查看可用工具
curl http://localhost:8080/api/v1/tools

# 多轮对话（使用 session ID）
$sessionId = (curl -X POST http://localhost:8080/api/v1/smart-chat `
  -H "Content-Type: application/json" `
  -d '{"message": "你好", "user_id": "test"}').session_id

curl -X POST http://localhost:8080/api/v1/smart-chat `
  -H "Content-Type: application/json" `
  -H "X-Session-ID: $sessionId" `
  -d '{"message": "刚才我们聊了什么？", "user_id": "test"}'
```

---

## 第二阶段：长期记忆 + 任务规划（2-3周）

### 需要添加的组件

1. **向量数据库集成** (Chroma/Pinecone)
2. **Embedding 模型** (OpenAI/BGE)
3. **任务规划器** (ReAct 模式)
4. **反思模块**

详细实现参考 [INTELLIGENT_AGENT_UPGRADE.md](INTELLIGENT_AGENT_UPGRADE.md)

---

## 📊 进度追踪

### Week 1 目标
- [ ] Redis 安装和配置
- [ ] 会话管理系统完成
- [ ] 工具注册中心完成
- [ ] 2个基础工具实现（计算器、日期）

### Week 2 目标
- [ ] 智能对话 API 完成
- [ ] 会话历史功能测试通过
- [ ] 工具调用流程测试通过
- [ ] 性能优化（缓存、连接池）

### Week 3-4 目标
- [ ] 向量数据库集成
- [ ] 长期记忆存储和检索
- [ ] 任务规划器实现
- [ ] 3+ 个高级工具（搜索、代码执行等）

---

## 🐛 常见问题

### Q1: Redis 连接失败
```bash
# 检查 Redis 是否运行
redis-cli ping
# 应该返回 PONG

# 检查防火墙
netstat -an | grep 6379
```

### Q2: 工具调用不生效
- 检查提示词中是否正确描述了工具
- 确保 LLM 温度参数不要太高（建议 0.7）
- 测试工具是否能独立运行

### Q3: 会话历史丢失
- 检查 Redis TTL 配置
- 确认会话 ID 正确传递
- 查看 Redis 中的键：`redis-cli keys "session:*"`

---

## 🎯 下一步

完成第一阶段后，继续实施：
1. 阅读 [INTELLIGENT_AGENT_UPGRADE.md](INTELLIGENT_AGENT_UPGRADE.md) 的第二、三阶段
2. 根据业务需求选择要实现的工具
3. 逐步添加高级功能（记忆、规划、反思）

有任何问题随时询问！🚀
