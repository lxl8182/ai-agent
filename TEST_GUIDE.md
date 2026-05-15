# 🧪 游戏攻略搜索功能 - 测试指南

## ⚠️ 当前状态

根据日志分析，目前存在以下问题：

### 1. Bing Search API Key 未配置
```
{"level":"warn","msg":"BING_SEARCH_API_KEY not set, web search will be disabled"}
```

**影响**: 
- ❌ 网络搜索功能被禁用
- ❌ 无法获取实时游戏攻略
- ✅ 但智能体仍可工作（使用 LLM 训练数据回答）

### 2. LLM API 连接问题
从历史日志看到有 API 连接错误，需要检查配置文件。

---

## 🔧 解决方案

### Step 1: 配置 LLM API

编辑 `configs/config.yaml`:

```yaml
agent:
  model: "qwen-turbo"
  api_key: "sk-your-actual-api-key"  # ← 填入你的阿里云 API Key
  base_url: "https://dashscope.aliyuncs.com/compatible-mode/v1"
  max_tokens: 2000
  temperature: 0.7
  system_prompt: |
    你是一个智能AI助手，可以访问网络搜索获取最新信息。
    
    当用户询问游戏攻略、实时新闻、最新数据等问题时，你会：
    1. 自动识别需要使用搜索工具
    2. 调用搜索工具获取最新信息
    3. 基于搜索结果给出准确、详细的回答
    
    请用中文回答用户的问题。
```

**获取 API Key**:
- 访问 https://dashscope.console.aliyun.com/
- 创建 API Key
- 复制并粘贴到配置文件

---

### Step 2: 配置 Bing Search API（可选，但推荐）

#### 方式 1: 环境变量（推荐）

**Windows PowerShell**:
```powershell
$env:BING_SEARCH_API_KEY="your-bing-api-key-here"
```

**永久设置**:
```powershell
[System.Environment]::SetEnvironmentVariable('BING_SEARCH_API_KEY', 'your-key', 'User')
```

**Linux/Mac**:
```bash
export BING_SEARCH_API_KEY="your-bing-api-key-here"
```

#### 方式 2: 获取 Bing API Key

1. 访问 https://portal.azure.com
2. 创建 "Bing Search v7" 资源
3. 获取 API Key
4. **免费额度**: 每月 1,000 次调用

---

### Step 3: 重启服务

```bash
# 停止当前服务（Ctrl+C）

# 重新启动
go run cmd/server/main.go
```

---

## 🧪 测试方法

### 方法 1: 使用 Web 界面（最简单）⭐

1. **打开浏览器**
   ```
   http://localhost:8080
   ```

2. **输入问题**
   - 例如："原神 雷电将军怎么培养"

3. **勾选选项**
   - ☑ 自动在浏览器中打开搜索结果
   - ☑ 使用智能搜索（推荐）

4. **点击搜索**
   - 查看结果

---

### 方法 2: 使用 PowerShell

```powershell
# 测试 1: 原神攻略
$body = @{message='原神 雷电将军培养'} | ConvertTo-Json
Invoke-RestMethod -Uri 'http://localhost:8080/api/v1/smart-chat' `
    -Method Post `
    -ContentType 'application/json; charset=utf-8' `
    -Body $body

# 测试 2: Boss 打法
$body = @{message='黑神话悟空 虎先锋怎么打'} | ConvertTo-Json
Invoke-RestMethod -Uri 'http://localhost:8080/api/v1/smart-chat' `
    -Method Post `
    -ContentType 'application/json; charset=utf-8' `
    -Body $body
```

---

### 方法 3: 使用 curl（Git Bash）

```bash
# 测试原神攻略
curl -X POST http://localhost:8080/api/v1/smart-chat \
  -H "Content-Type: application/json" \
  -d "{\"message\": \"原神 雷电将军培养\"}"

# 测试 Boss 打法
curl -X POST http://localhost:8080/api/v1/smart-chat \
  -H "Content-Type: application/json" \
  -d "{\"message\": \"黑神话悟空 虎先锋怎么打\"}"
```

---

## 📊 预期结果

### 情况 1: 只配置了 LLM API（无 Bing）

**行为**:
- ✅ 智能体可以回答问题
- ❌ 不会调用搜索工具
- ⚠️ 回答基于训练数据（可能不是最新的）

**示例响应**:
```json
{
  "success": true,
  "data": "雷电将军是原神中的5星雷元素角色...\n\n注意：由于未配置搜索工具，以上信息可能不是最新的。"
}
```

---

### 情况 2: 同时配置了 LLM + Bing API

**行为**:
- ✅ 智能识别游戏问题
- ✅ 调用 Bing 搜索最新攻略
- ✅ 整合搜索结果生成回答
- ✅ 自动打开浏览器（如果勾选）

**示例响应**:
```json
{
  "success": true,
  "data": "🎮 **原神 - 雷电将军培养** 攻略搜索结果:\n\n找到 8 个相关结果:\n\n1. **雷电将军全面培养指南**\n   主词条推荐：...\n   链接: https://...\n\n..."
}
```

---

## 🐛 常见问题排查

### Q1: 返回乱码或无法识别的字符

**原因**: 编码问题或 API 配置错误

**解决**:
1. 检查 `config.yaml` 中的 API Key 是否正确
2. 确保使用 UTF-8 编码
3. 查看日志文件 `logs/agent.log`

---

### Q2: 提示 "BING_SEARCH_API_KEY not set"

**这是正常的**，如果不配置 Bing API：
- 搜索功能会被禁用
- 但仍可以使用 LLM 回答

**如果想启用搜索**:
- 按照上面的步骤配置 Bing API Key

---

### Q3: API 连接超时

**可能原因**:
1. 网络连接问题
2. API Key 无效
3. 防火墙阻止

**解决**:
1. 检查网络连接
2. 验证 API Key 是否有效
3. 检查防火墙设置

---

### Q4: 端口 8080 被占用

**解决**:
```powershell
# 查找占用端口的进程
netstat -ano | findstr :8080

# 结束进程（替换 PID）
taskkill /F /PID <PID>

# 或者修改端口
# 编辑 configs/config.yaml
server:
  port: 8081  # 改为其他端口
```

---

## 📝 快速测试清单

### 最小化测试（只需 LLM API）

- [ ] 配置 `configs/config.yaml` 中的 `api_key`
- [ ] 启动服务 `go run cmd/server/main.go`
- [ ] 访问 http://localhost:8080
- [ ] 输入问题并测试

**预期**: 能收到回答（可能不是最新的）

---

### 完整测试（LLM + Bing API）

- [ ] 配置 LLM API Key
- [ ] 配置 Bing API Key（环境变量）
- [ ] 重启服务
- [ ] 访问 Web 界面
- [ ] 勾选"自动打开浏览器"
- [ ] 搜索游戏攻略

**预期**: 
- ✅ 页面显示智能回答
- ✅ 浏览器自动打开搜索结果
- ✅ 回答包含最新信息

---

## 🎯 推荐的测试问题

### 简单测试
1. "你好"
2. "什么是人工智能"

### 游戏相关测试
1. "原神 雷电将军培养"
2. "黑神话悟空 虎先锋打法"
3. "塞尔达传说 神庙位置"
4. "王者荣耀 S35赛季英雄"

### 边界测试
1. "" （空输入）
2. "asdfghjkl" （无意义字符）
3. 非常长的问题

---

## 📈 性能指标

### 响应时间

| 配置 | 平均时间 | 说明 |
|------|---------|------|
| 仅 LLM | 1-2秒 | 直接调用 API |
| LLM + Bing | 3-5秒 | 搜索(1-2s) + LLM(2-3s) |

### 准确率

| 指标 | 目标值 |
|------|--------|
| 意图识别 | > 90% |
| 搜索相关性 | > 80% |
| 回答质量 | > 85% |

---

## 🔍 调试技巧

### 1. 查看日志

```bash
# 实时查看日志
Get-Content logs/agent.log -Wait

# 或
tail -f logs/agent.log
```

### 2. 检查请求

浏览器开发者工具 → Network 标签 → 查看请求和响应

### 3. 测试 API 连通性

```powershell
# 测试健康检查
Invoke-RestMethod -Uri 'http://localhost:8080/health'

# 应该返回
# status: ok
# time: 2026-05-15T...
```

---

## 💡 最佳实践

### 1. 先测试 LLM API

确保基础对话功能正常，再添加搜索功能。

### 2. 逐步测试

1. 健康检查 → `/health`
2. 简单对话 → `/api/v1/simple-chat`
3. 智能对话 → `/api/v1/smart-chat`

### 3. 使用 Web 界面

Web 界面最直观，适合快速测试。

### 4. 记录测试结果

保存成功的响应，用于对比和优化。

---

## 🎉 开始测试！

### 立即行动

1. **配置 API Keys**
   ```yaml
   # configs/config.yaml
   agent:
     api_key: "sk-your-key"
   ```

2. **启动服务**
   ```bash
   go run cmd/server/main.go
   ```

3. **打开浏览器**
   ```
   http://localhost:8080
   ```

4. **测试搜索**
   - 输入："原神 雷电将军培养"
   - 点击搜索
   - 查看结果

---

## 📞 需要帮助？

如有问题：

1. 📖 查看日志 `logs/agent.log`
2. 💬 检查配置文件 `configs/config.yaml`
3. 🔍 确认 API Keys 有效
4. 📧 联系技术支持

---

**祝你测试顺利！** 🎮✨
