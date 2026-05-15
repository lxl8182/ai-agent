# 🎮 AI 游戏攻略助手

一个基于 Go 语言的智能游戏攻略助手，能够自动检测游戏问题并打开浏览器搜索最新攻略，同时提供 AI 智能对话功能。

## ✨ 核心功能

- 🤖 **AI 智能对话** - 接入 LLM（通义千问等），支持多轮对话
- 🎮 **游戏智能检测** - 自动识别 60+ 款热门游戏
- 🌐 **自动浏览器搜索** - 检测到游戏问题时自动打开浏览器搜索
- 💾 **对话历史记忆** - 自动保存对话记录，刷新页面不丢失
- 📝 **Markdown 渲染** - 美观的格式化输出
- ⚙️ **游戏列表管理** - 可实时添加/删除游戏，无需修改代码

## 🚀 快速开始

### 1. 环境要求

- Go 1.21+
- LLM API Key（阿里云通义千问等）

### 2. 配置 API 密钥

编辑 `configs/config.yaml`：

```yaml
agent:
  model: "qwen-turbo"
  api_key: "sk-your-api-key-here"  # ← 填入你的 API Key
  base_url: "https://dashscope.aliyuncs.com/compatible-mode/v1"
```

**获取 API Key**:
- 阿里云通义千问: https://dashscope.console.aliyun.com/

### 3. 安装依赖

```bash
go mod tidy
```

### 4. 运行服务

```bash
go run cmd/server/main.go
```

服务将在 `http://localhost:8080` 启动

## 🌐 访问前端页面

打开浏览器访问：

```
http://localhost:8080/agent
```

## 💡 使用示例

### 游戏问题（自动打开浏览器）

**输入**: "原神 雷电将军怎么培养"

**效果**:
1. 🌐 自动打开新标签页，显示 Bing 搜索结果
2. 💬 AI 助手提供基于训练数据的建议

### 普通问题

**输入**: "什么是人工智能"

**效果**:
- 💬 AI 直接回答
- ❌ 不打开浏览器

## 📚 相关文档

- 📘 [使用说明.md](使用说明.md) - 详细使用指南
- 📗 [完整AI助手指南.md](完整AI助手指南.md) - 功能详细说明
- 📙 [架构设计.md](架构设计.md) - 技术架构说明
- 📕 [API使用示例.md](API使用示例.md) - API 调用示例
- 📒 [更新日志.md](更新日志.md) - 版本更新记录

## 🛠️ 技术栈

- **后端**: Go 1.21 + Gin Framework
- **日志**: Zap 结构化日志
- **前端**: HTML5 + JavaScript + Marked.js
- **配置**: YAML
- **LLM**: 阿里云通义千问（可替换）

## 📁 项目结构

```
ai-agent/
├── cmd/server/          # 主程序入口
├── configs/             # 配置文件
├── internal/            # 内部包
│   ├── agent/          # Agent核心逻辑
│   ├── config/         # 配置管理
│   └── handler/        # HTTP处理器
├── pkg/logger/         # 日志模块
├── web/                # 前端页面
│   └── full-agent.html # 唯一前端页面
└── README.md           # 项目说明
```

## 🔧 扩展开发

### 添加新的游戏

在页面中点击“🎮 管理游戏列表”按钮，即可实时添加/删除游戏，无需修改代码。

### 添加新的 API 端点

在 `internal/handler/handler.go` 中添加新的处理函数和路由。

## 📝 许可证

MIT License

---

**祝你使用愉快！** 🎮✨
