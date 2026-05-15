# GitHub 上传指南

本文档帮助你安全地将项目上传到GitHub，同时保护API密钥等敏感信息。

## ⚠️ 重要安全提示

**绝对不要**将以下内容提交到GitHub：
- ❌ API密钥（`configs/config.yaml` 中的 `api_key`）
- ❌ `.env` 文件
- ❌ 任何包含密码、令牌的文件
- ❌ 日志文件（可能包含敏感信息）

## 📋 上传前检查清单

### 1. 确认 .gitignore 已配置
以下文件已被排除，不会上传：
- ✅ `configs/config.yaml` - 包含API密钥的配置文件
- ✅ `.env` - 环境变量文件
- ✅ `logs/` - 日志目录
- ✅ `build/` - 编译输出
- ✅ `*.exe` - 可执行文件

### 2. 准备安全的配置文件
项目中已包含 `configs/config.yaml.example`，其他人可以复制它并填入自己的API密钥。

### 3. 清理敏感信息
如果之前已经提交了包含API密钥的文件，需要：
```bash
# 从Git历史中移除敏感文件
git rm --cached configs/config.yaml
git commit -m "Remove sensitive config file"
```

## 🚀 上传步骤

### 方法1：使用命令行（推荐）

#### 1. 初始化Git仓库（如果还没有）
```bash
cd E:\job\code\ai-agent
git init
```

#### 2. 添加所有文件
```bash
git add .
```

#### 3. 检查将要提交的文件
```bash
git status
```
**确认没有以下文件：**
- `configs/config.yaml`
- `.env`
- `logs/` 目录下的文件

#### 4. 提交代码
```bash
git commit -m "Initial commit: AI Agent project with Alibaba Cloud integration"
```

#### 5. 在GitHub上创建新仓库
- 访问 https://github.com/new
- 仓库名：`ai-agent`（或其他你喜欢的名字）
- 设置为 Public（公开）或 Private（私有）
- **不要**勾选 "Initialize with README"（因为我们已有README）

#### 6. 关联远程仓库
```bash
git remote add origin https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
```
替换 `YOUR_USERNAME` 和 `YOUR_REPO_NAME` 为你的实际信息。

#### 7. 推送到GitHub
```bash
git branch -M main
git push -u origin main
```

### 方法2：使用GitHub Desktop

1. 打开 GitHub Desktop
2. File → Add Local Repository → 选择 `E:\job\code\ai-agent`
3. 填写 Commit 信息
4. Publish repository
5. 设置为 Public 或 Private

## 🔒 安全最佳实践

### 1. 使用环境变量（推荐）
创建 `.env` 文件（不会被提交）：
```env
AGENT_API_KEY=sk-your-actual-api-key
AGENT_MODEL=qwen-turbo
```

### 2. 定期轮换API密钥
- 每3-6个月更换一次API密钥
- 如果怀疑泄露，立即在阿里云控制台撤销旧密钥

### 3. 使用GitHub Secrets（用于CI/CD）
如果在GitHub Actions中使用：
1. 进入仓库 Settings → Secrets and variables → Actions
2. 点击 "New repository secret"
3. 添加 `AGENT_API_KEY` 等敏感信息

### 4. 设置仓库可见性
- **Private仓库**：只有你和协作者可以访问
- **Public仓库**：任何人都可以看到代码（但不要包含密钥）

## ✅ 验证上传成功

### 检查GitHub仓库
1. 访问你的GitHub仓库页面
2. 确认文件列表中包含：
   - ✅ `README.md`
   - ✅ `cmd/server/main.go`
   - ✅ `internal/` 目录
   - ✅ `configs/config.yaml.example`
   - ❌ **不应该有** `configs/config.yaml`
   - ❌ **不应该有** `.env`

### 测试克隆
```bash
# 在新目录测试
git clone https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git test-clone
cd test-clone
ls configs/
# 应该只看到 config.yaml.example
```

## 📝 首次使用者说明

在README中添加以下说明，帮助其他人使用你的项目：

```markdown
## 首次使用

1. 复制配置模板：
   ```bash
   cp configs/config.yaml.example configs/config.yaml
   ```

2. 编辑 `configs/config.yaml`，填入你的API密钥：
   ```yaml
   agent:
     api_key: "your-api-key-here"
   ```

3. 或者使用环境变量（推荐）：
   ```bash
   export AGENT_API_KEY="your-api-key-here"
   ```
```

## 🐛 常见问题

### Q: 不小心提交了API密钥怎么办？
A: 
1. 立即在阿里云控制台撤销该密钥
2. 生成新密钥
3. 使用 `git filter-branch` 或 BFG Repo-Cleaner 清理历史
4. 或者删除仓库重新创建

### Q: 如何查看哪些文件会被提交？
A:
```bash
git status
git diff --cached --name-only
```

### Q: 如何从Git历史中彻底删除文件？
A:
```bash
# 使用 git filter-branch
git filter-branch --force --index-filter \
  'git rm --cached --ignore-unmatch configs/config.yaml' \
  --prune-empty HEAD

# 强制推送（谨慎使用）
git push origin --force --all
```

### Q: Private仓库是否安全？
A: 
- Private仓库比Public更安全
- 但仍建议不要提交敏感信息
- 只有受信任的协作者才能访问

## 📞 需要帮助？

- GitHub文档：https://docs.github.com/
- Git手册：https://git-scm.com/doc
- 阿里云API管理：https://dashscope.console.aliyun.com/

---

**记住：安全第一！永远不要在代码中硬编码API密钥。** 🔐
