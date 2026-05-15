# GitHub 上传步骤

## 当前状态
- ✅ 代码已提交到本地Git
- ✅ 远程仓库已配置：https://github.com/lxl8182/ai-agent.git
- ⏳ 等待推送到GitHub

## 推送步骤

### 方法1：使用命令行（推荐）

打开PowerShell或CMD，运行：

```bash
cd E:\job\code\ai-agent
git push -u origin main
```

首次推送时，系统会要求你登录GitHub账号进行认证。

### 方法2：使用脚本

双击运行项目根目录下的：
```
push_to_github.bat
```

### 方法3：使用GitHub Desktop

1. 打开 GitHub Desktop
2. File → Add Local Repository
3. 选择 `E:\job\code\ai-agent`
4. 点击 "Publish repository"

## 如果遇到问题

### 问题1：Repository not found
**原因：** GitHub上还没有创建仓库

**解决：**
1. 访问 https://github.com/new
2. 创建名为 `ai-agent` 的仓库
3. ❌ 不要勾选任何初始化选项
4. 创建完成后再次推送

### 问题2：网络连接失败
**原因：** 网络不稳定或被墙

**解决：**
1. 检查网络连接
2. 如果使用代理，确保Git配置了代理：
   ```bash
   git config --global http.proxy http://127.0.0.1:7890
   git config --global https.proxy http://127.0.0.1:7890
   ```
3. 稍后重试

### 问题3：认证失败
**原因：** 需要GitHub认证

**解决：**
1. 浏览器会自动打开让你登录
2. 或使用Personal Access Token：
   - 访问：https://github.com/settings/tokens
   - 生成新token（勾选repo权限）
   - 推送时使用token作为密码

## 验证上传成功

访问：https://github.com/lxl8182/ai-agent

确认文件列表中包含：
- ✅ README.md
- ✅ cmd/server/main.go
- ✅ internal/ 目录
- ✅ configs/config.yaml.example
- ❌ **不应该有** configs/config.yaml

## 安全提醒

✅ 已保护的文件（不会上传）：
- configs/config.yaml（包含API密钥）
- .env
- logs/
- build/

---

**准备好后，运行 `git push -u origin main` 即可！** 🚀
