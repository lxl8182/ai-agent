# 🚀 快速上传到GitHub

你的代码已经准备好上传了！按照以下步骤操作：

## ✅ 已完成的工作

- ✅ Git仓库已初始化
- ✅ 代码已提交（25个文件，3180行）
- ✅ API密钥已被排除（configs/config.yaml 不在Git中）
- ✅ 配置模板已创建（configs/config.yaml.example）

## 📤 上传步骤

### 第1步：在GitHub上创建仓库

1. 访问 https://github.com/new
2. 填写信息：
   - **Repository name**: `ai-agent` （或你喜欢的名字）
   - **Description**: `AI Agent service with Alibaba Cloud Qwen integration`
   - **Public/Private**: 选择 Public（公开）或 Private（私有）
   - ❌ **不要勾选** "Initialize with README"
   - ❌ **不要勾选** "Add .gitignore"
   - ❌ **不要勾选** "Choose a license"
3. 点击 "Create repository"

### 第2步：关联远程仓库并推送

复制GitHub显示的命令，或者运行：

```bash
cd E:\job\code\ai-agent

# 关联远程仓库（替换 YOUR_USERNAME 和 YOUR_REPO_NAME）
git remote add origin https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git

# 重命名分支为 main
git branch -M main

# 推送到GitHub
git push -u origin main
```

**示例：**
```bash
git remote add origin https://github.com/john-doe/ai-agent.git
git branch -M main
git push -u origin main
```

### 第3步：验证上传

访问你的GitHub仓库页面，确认看到以下文件：
- ✅ README.md
- ✅ cmd/server/main.go
- ✅ internal/ 目录
- ✅ configs/config.yaml.example
- ❌ **不应该有** configs/config.yaml
- ❌ **不应该有** .env

## 🔐 安全确认

### 已保护的文件
以下文件**不会**被上传到GitHub：
- `configs/config.yaml` - 包含你的真实API密钥
- `.env` - 环境变量文件
- `logs/` - 日志目录
- `build/` - 编译输出
- `*.exe` - 可执行文件

### 其他人如何使用

其他人克隆你的仓库后，需要：

1. 复制配置模板：
   ```bash
   cp configs/config.yaml.example configs/config.yaml
   ```

2. 编辑 `configs/config.yaml`，填入他们自己的API密钥

3. 或者设置环境变量：
   ```bash
   export AGENT_API_KEY="their-api-key"
   ```

## 📝 可选：添加LICENSE

如果想添加开源许可证：

```bash
# 添加MIT许可证
curl -o LICENSE https://raw.githubusercontent.com/github/gitignore/master/MIT.LICENSE
git add LICENSE
git commit -m "Add MIT license"
git push
```

或在GitHub网页界面添加：
1. 进入仓库 → Add file → Create new file
2. 文件名输入 `LICENSE`
3. 点击 "Choose a license template"
4. 选择 MIT License
5. 点击 "Commit changes"

## 🎯 后续更新代码

当你修改代码后，推送更新：

```bash
git add .
git commit -m "描述你的修改"
git push
```

## 💡 提示

### 使用SSH而不是HTTPS（可选）
如果配置了SSH密钥：
```bash
git remote set-url origin git@github.com:YOUR_USERNAME/YOUR_REPO_NAME.git
```

### 查看远程仓库
```bash
git remote -v
```

### 拉取最新代码
```bash
git pull origin main
```

## ❓ 遇到问题？

### 问题1：提示需要登录
```bash
# 使用Personal Access Token
# 1. GitHub → Settings → Developer settings → Personal access tokens
# 2. 生成新token（勾选 repo 权限）
# 3. 推送时使用token作为密码
```

### 问题2：推送失败
```bash
# 检查远程仓库地址
git remote -v

# 重新设置
git remote set-url origin https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
```

### 问题3：想改为Private仓库
1. GitHub仓库 → Settings
2. Danger Zone → Change visibility
3. 选择 Make private

## 🎊 完成！

上传成功后，你可以：
- 分享仓库链接给他人
- 在README中添加badge
- 设置GitHub Actions CI/CD
- 邀请协作者

---

**记住：永远不要在代码中硬编码API密钥！** 🔐

祝你好运！🚀
