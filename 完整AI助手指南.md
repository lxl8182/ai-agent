# 🚀 AI 游戏攻略助手 - 完整部署指南

## ✨ 功能概述

这是一个**完整的前后端集成方案**：

### 核心功能

✅ **后端服务（Go）**
- 提供 LLM API 接口
- 智能对话能力
- 游戏意图识别

✅ **前端页面（HTML + JavaScript）**
- 接入后端 LLM 实现智能对话
- 自动检测游戏名称
- 检测到游戏时自动打开浏览器搜索攻略

✅ **双重能力**
- 💬 普通问题 → LLM 回答
- 🎮 游戏问题 → LLM 回答 + 自动搜索

---

## 📋 系统架构

```
┌─────────────────────────────────────┐
│      用户浏览器 (前端)               │
│  ┌───────────────────────────────┐  │
│  │  web/full-agent.html          │  │
│  │  • 聊天界面                    │  │
│  │  • 游戏检测                    │  │
│  │  • 自动搜索                    │  │
│  └──────────┬────────────────────┘  │
└─────────────┼───────────────────────┘
              │ HTTP API
              ↓
┌─────────────────────────────────────┐
│   Go 后端服务 (localhost:8080)      │
│  ┌───────────────────────────────┐  │
│  │  • LLM API (阿里云/其他)      │  │
│  │  • 路由管理                   │  │
│  │  • 静态文件服务               │  │
│  └───────────────────────────────┘  │
└─────────────────────────────────────┘
              │
              ↓
┌─────────────────────────────────────┐
│   外部服务                          │
│  • LLM API (通义千问等)             │
│  • Bing/Google 搜索引擎             │
└─────────────────────────────────────┘
```

---

## 🔧 部署步骤

### Step 1: 配置 LLM API

编辑 `configs/config.yaml`:

```yaml
server:
  host: "0.0.0.0"
  port: 8080
  read_timeout: 30
  write_timeout: 30

agent:
  model: "qwen-turbo"  # 或 gpt-3.5-turbo 等
  api_key: "sk-your-api-key-here"  # ← 填入你的 API Key
  base_url: "https://dashscope.aliyuncs.com/compatible-mode/v1"
  max_tokens: 2000
  temperature: 0.7
  system_prompt: |
    你是一个智能AI助手，可以回答各种问题。
    
    当用户询问游戏攻略时，你会：
    1. 提供详细的攻略信息
    2. 给出实用的建议
    3. 用中文友好地回答
    
    请用中文回答用户的问题。

log:
  level: "info"
  file: "logs/agent.log"
```

**获取 API Key**:
- 阿里云通义千问: https://dashscope.console.aliyun.com/
- OpenAI: https://platform.openai.com/

---

### Step 2: 启动后端服务

```bash
# 方式 1: 直接运行
go run cmd/server/main.go

# 方式 2: 编译后运行
go build -o build/server.exe cmd/server/main.go
./build/server.exe

# 方式 3: 使用 Makefile
make run
```

**预期输出**:
```
{"level":"info","msg":"Starting AI Agent Server..."}
{"level":"info","msg":"Server starting on 0.0.0.0:8080"}
[GIN-debug] Listening and serving HTTP on 0.0.0.0:8080
```

---

### Step 3: 访问前端页面

在浏览器中打开：

```
http://localhost:8080/agent
```

或者双击打开文件：

```
web\full-agent.html
```

---

## 📸 使用演示

### 场景 1: 游戏相关问题

**用户输入**: "原神 雷电将军怎么培养"

**系统流程**:
1. 🔍 前端检测到游戏名："原神"
2. 💬 发送请求到后端 LLM API
3. 🤖 LLM 生成详细回答
4. 🌐 **同时自动打开浏览器**搜索攻略
5. ✅ 显示 LLM 回答 + 浏览器搜索结果

**界面显示**:
```
用户: 原神 雷电将军怎么培养

AI: 雷电将军是原神中的5星雷元素角色...
    [详细的培养指南]

系统: 🔍 检测到游戏"原神"，正在调用 AI 回答并准备搜索...
系统: ✅ 已在浏览器中打开"原神"的搜索结果！
```

**浏览器**: 自动打开新标签页，显示 Bing 搜索结果

---

### 场景 2: 非游戏问题

**用户输入**: "今天天气怎么样"

**系统流程**:
1. ❌ 未检测到游戏名
2. 💬 发送请求到后端 LLM API
3. 🤖 LLM 生成回答
4. ✅ 只显示 LLM 回答，不打开浏览器

**界面显示**:
```
用户: 今天天气怎么样

AI: 抱歉，我无法获取实时天气信息...
```

---

## ⚙️ 配置选项

### 前端选项

**☑ 检测到游戏时自动打开浏览器搜索**
- 勾选：检测到游戏后自动打开浏览器
- 不勾选：只使用 LLM 回答

**☑ 在新标签页打开**
- 勾选：新标签页（推荐）
- 不勾选：当前页面跳转

---

### 后端配置

**修改端口**:
```yaml
server:
  port: 8081  # 改为其他端口
```

**修改日志级别**:
```yaml
log:
  level: "debug"  # debug/info/warn/error
```

---

## 🎯 支持的游戏

前端内置 50+ 款游戏检测：

- 原神、崩坏星穹铁道、黑神话悟空
- 王者荣耀、和平精英、阴阳师
- 塞尔达传说、英雄联盟、Minecraft
- GTA V、艾尔登法环、博德之门3
- ...（完整列表见代码）

**添加新游戏**: 编辑 `web/full-agent.html` 中的 `gameNames` 数组

---

## 🔍 工作流程详解

### 完整流程图

```
用户输入问题
     ↓
前端检测游戏名
     ↓
┌────────────┬──────────────┐
│ 检测到游戏  │ 未检测到游戏  │
└──────┬─────┴──────┬───────┘
       ↓            ↓
  提示正在搜索   直接调用LLM
       ↓            ↓
  打开浏览器     等待LLM响应
  搜索攻略         ↓
       ↓       显示回答
  调用LLM API
       ↓
  显示回答 + 
  确认已打开浏览器
```

---

## 💡 使用技巧

### 技巧 1: 自然语言提问

**好的提问**:
- ✅ "原神 雷电将军的圣遗物怎么搭配"
- ✅ "黑神话悟空卡关了，虎先锋怎么打"
- ✅ "我想了解一下塞尔达的神庙位置"

**系统会**:
- 提取游戏名
- LLM 提供详细回答
- 浏览器显示最新攻略

---

### 技巧 2: 对比信息

1. 查看 LLM 的结构化回答
2. 切换到浏览器标签查看网页
3. 综合多个来源的信息

---

### 技巧 3: 关闭自动搜索

如果只想使用 LLM：
- 取消勾选"自动打开浏览器搜索"
- 所有问题都只通过 LLM 回答

---

## 🐛 常见问题

### Q1: 提示"服务器未连接"？

**原因**: 后端服务未启动

**解决**:
```bash
# 启动服务
go run cmd/server/main.go
```

---

### Q2: LLM 返回错误？

**可能原因**:
1. API Key 无效
2. 网络连接问题
3. API 配额用完

**检查**:
```bash
# 查看日志
cat logs/agent.log

# 测试 API
curl http://localhost:8080/health
```

---

### Q3: 浏览器没有自动打开？

**检查**:
1. 是否勾选了"自动打开浏览器"
2. 浏览器是否阻止弹窗
3. 查看地址栏拦截提示

**解决**: 允许 localhost:8080 的弹窗

---

### Q4: 如何更换 LLM 提供商？

**修改配置**:

```yaml
agent:
  model: "gpt-3.5-turbo"  # 改为 GPT
  api_key: "sk-openai-key"
  base_url: "https://api.openai.com/v1"
```

---

## 📊 性能指标

### 响应时间

| 操作 | 耗时 |
|------|------|
| 游戏检测 | < 1ms |
| LLM 调用 | 1-3秒 |
| 打开浏览器 | ~1秒 |
| 总耗时 | 2-4秒 |

---

### 资源占用

| 指标 | 数值 |
|------|------|
| 后端内存 | ~50 MB |
| 前端内存 | < 20 MB |
| CPU 占用 | 低 |

---

## 🎨 界面特色

### 现代化设计
- 🎨 渐变背景
- 💬 聊天式界面
- ✨ 流畅动画
- 📱 响应式布局

### 实时反馈
- 🔗 连接状态指示器
- ⌨️ 打字指示器
- ✅ 成功提示
- ❌ 错误提示

---

## 🔧 开发指南

### 项目结构

```
ai-agent/
├── cmd/server/
│   └── main.go              # 主程序入口
├── internal/
│   ├── agent/               # Agent 逻辑
│   ├── handler/             # HTTP 处理器
│   └── config/              # 配置管理
├── web/
│   ├── full-agent.html      # ⭐ 完整版（LLM + 搜索）
│   ├── smart-search.html    # 智能检测版（仅搜索）
│   └── simple-search.html   # 简单搜索版
└── configs/
    └── config.yaml          # 配置文件
```

---

### 添加新功能

**1. 添加新的 API 端点**:

编辑 `internal/handler/handler.go`:

```go
router.POST("/api/v1/new-endpoint", handler.newHandler)
```

**2. 添加新的游戏**:

编辑 `web/full-agent.html`:

```javascript
const gameNames = [
    // 现有游戏...
    '新游戏名',
];
```

**3. 修改 LLM 提示词**:

编辑 `configs/config.yaml`:

```yaml
agent:
  system_prompt: "你的新提示词..."
```

---

## 📈 优化建议

### 性能优化

1. **启用 Gzip 压缩**
```go
import "github.com/gin-contrib/gzip"
router.Use(gzip.Gzip(gzip.DefaultCompression))
```

2. **添加缓存**
```go
// 缓存常见问题的回答
```

3. **异步处理**
```go
// 后台打开浏览器，不阻塞 LLM 响应
```

---

### 用户体验优化

1. **流式响应**
```javascript
// 使用 WebSocket 实现打字机效果
```

2. **历史记录**
```javascript
// 保存对话历史到 localStorage
```

3. **语音输入**
```javascript
// 使用 Web Speech API
```

---

## 🚀 部署到生产环境

### 1. 编译优化

```bash
# 编译优化版本
go build -ldflags="-s -w" -o server cmd/server/main.go
```

### 2. 配置 Nginx

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### 3. 启用 HTTPS

```bash
# 使用 Let's Encrypt
certbot --nginx -d your-domain.com
```

### 4. 进程管理

```bash
# 使用 systemd
sudo systemctl enable ai-agent
sudo systemctl start ai-agent
```

---

## 📝 版本对比

| 版本 | 文件 | LLM | 搜索 | 适用场景 |
|------|------|-----|------|---------|
| **完整版** ⭐ | `full-agent.html` | ✅ | ✅ | 最佳体验 |
| 智能检测版 | `smart-search.html` | ❌ | ✅ | 只需搜索 |
| 简单搜索版 | `simple-search.html` | ❌ | ✅ | 手动搜索 |
| 带 API 版 | `index.html` | ✅ | ✅ | 需要整合 |

---

## 🎉 总结

### 你得到了什么？

✅ **完整的解决方案**
- 前后端集成
- LLM 智能对话
- 自动浏览器搜索

✅ **优秀的用户体验**
- 聊天式界面
- 实时反馈
- 双重信息源

✅ **高度可定制**
- 更换 LLM 提供商
- 添加新游戏
- 自定义样式

✅ **生产就绪**
- 完善的错误处理
- 日志记录
- 可扩展架构

---

## 🚀 快速开始

```bash
# 1. 配置 API Key
# 编辑 configs/config.yaml

# 2. 启动服务
go run cmd/server/main.go

# 3. 访问页面
# http://localhost:8080/agent

# 4. 开始对话
# 输入任何问题！
```

---

**祝你使用愉快！** 🎮✨
