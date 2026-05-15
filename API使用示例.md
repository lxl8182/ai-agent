# API使用示例

## 1. 健康检查

```bash
curl http://localhost:8080/health
```

响应：
```json
{
  "status": "ok",
  "time": "2024-01-01T12:00:00Z"
}
```

## 2. 简单聊天

```bash
curl -X POST http://localhost:8080/api/v1/simple-chat \
  -H "Content-Type: application/json" \
  -d '{
    "message": "你好，请介绍一下你自己"
  }'
```

响应：
```json
{
  "success": true,
  "data": "你好！我是一个AI助手..."
}
```

## 3. 高级聊天

```bash
curl -X POST http://localhost:8080/api/v1/chat \
  -H "Content-Type: application/json" \
  -d '{
    "message": "什么是人工智能？"
  }'
```

响应：
```json
{
  "success": true,
  "data": "人工智能（Artificial Intelligence，简称AI）..."
}
```

## 4. 错误处理

如果请求格式不正确：

```bash
curl -X POST http://localhost:8080/api/v1/simple-chat \
  -H "Content-Type: application/json" \
  -d '{
    "wrong_field": "test"
  }'
```

响应：
```json
{
  "success": false,
  "error": "Invalid request: Key: 'ChatRequest.Message' Error:Field validation for 'Message' failed on the 'required' tag"
}
```

## 5. 使用Python调用示例

```python
import requests
import json

# 简单聊天
url = "http://localhost:8080/api/v1/simple-chat"
payload = {
    "message": "请写一首关于春天的诗"
}

response = requests.post(url, json=payload)
print(response.json())

# 检查响应
if response.status_code == 200:
    data = response.json()
    if data['success']:
        print("AI回复:", data['data'])
    else:
        print("错误:", data['error'])
else:
    print("HTTP错误:", response.status_code)
```

## 6. 使用JavaScript调用示例

```javascript
async function chatWithAI(message) {
    const url = 'http://localhost:8080/api/v1/simple-chat';
    
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ message: message })
        });
        
        const data = await response.json();
        
        if (data.success) {
            console.log('AI回复:', data.data);
            return data.data;
        } else {
            console.error('错误:', data.error);
            return null;
        }
    } catch (error) {
        console.error('请求失败:', error);
        return null;
    }
}

// 使用示例
chatWithAI('请解释一下量子计算');
```

## 7. 配置不同的模型

编辑 `configs/config.yaml` 来切换不同的模型：

```yaml
agent:
  # OpenAI GPT-4
  model: "gpt-4"
  api_key: "your-openai-key"
  base_url: "https://api.openai.com/v1"
  
  # 或者使用其他兼容的API
  # model: "claude-2"
  # base_url: "https://api.anthropic.com/v1"
```

## 8. 多轮对话示例

虽然当前API是单轮的，但你可以在客户端维护对话历史：

```python
import requests

class ChatSession:
    def __init__(self):
        self.url = "http://localhost:8080/api/v1/simple-chat"
        self.history = []
    
    def send_message(self, message):
        # 添加用户消息到历史
        self.history.append(f"User: {message}")
        
        # 发送请求
        response = requests.post(self.url, json={"message": message})
        data = response.json()
        
        if data['success']:
            ai_response = data['data']
            # 添加AI回复到历史
            self.history.append(f"AI: {ai_response}")
            return ai_response
        else:
            return f"Error: {data['error']}"
    
    def get_history(self):
        return "\n".join(self.history)

# 使用示例
session = ChatSession()
print(session.send_message("你好"))
print(session.send_message("你能帮我做什么？"))
print(session.get_history())
```

## 9. 性能测试

使用ab或wrk进行压力测试：

```bash
# 使用ab测试
ab -n 100 -c 10 -p payload.json -T application/json \
   http://localhost:8080/api/v1/simple-chat

# payload.json 内容:
# {"message": "test message"}
```

## 10. 环境变量配置（可选）

你也可以通过环境变量覆盖配置文件：

```bash
export AGENT_API_KEY="your-api-key"
export AGENT_MODEL="gpt-4"
export SERVER_PORT=9000

go run cmd/server/main.go
```
