package agent

import (
	"ai-agent/internal/config"
	"ai-agent/internal/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type AgentService struct {
	config       config.AgentConfig
	toolRegistry *tools.Registry
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
	Temperature float64       `json:"temperature,omitempty"`
}

type ChatResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func NewAgentService(cfg config.AgentConfig) *AgentService {
	return &AgentService{
		config: cfg,
	}
}

// NewAgentServiceWithTools 创建带工具支持的 Agent 服务
func NewAgentServiceWithTools(cfg config.AgentConfig, registry *tools.Registry) *AgentService {
	return &AgentService{
		config:       cfg,
		toolRegistry: registry,
	}
}

func (a *AgentService) Chat(messages []ChatMessage) (string, error) {
	reqBody := ChatRequest{
		Model:       a.config.Model,
		Messages:    messages,
		MaxTokens:   a.config.MaxTokens,
		Temperature: a.config.Temperature,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", a.config.BaseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.config.APIKey)
	// 阿里云DashScope可能需要额外的header
	if strings.Contains(a.config.BaseURL, "dashscope") {
		req.Header.Set("X-DashScope-SSE", "disable")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return chatResp.Choices[0].Message.Content, nil
}

func (a *AgentService) SimpleChat(userMessage string) (string, error) {
	messages := []ChatMessage{
		{Role: "system", Content: a.config.SystemPrompt},
		{Role: "user", Content: userMessage},
	}

	return a.Chat(messages)
}

// SmartChat 智能对话（支持工具调用）
func (a *AgentService) SmartChat(userMessage string) (string, error) {
	if a.toolRegistry == nil {
		// 如果没有工具注册中心，退化为普通对话
		return a.SimpleChat(userMessage)
	}

	// 1. 首先判断是否需要调用工具
	shouldUseTool, toolName, toolArgs := a.analyzeIntent(userMessage)

	if shouldUseTool && toolName != "" {
		// 2. 执行工具调用
		toolResult, err := a.toolRegistry.Execute(toolName, toolArgs)
		if err != nil {
			// 工具调用失败，记录错误但仍尝试回答
			fallbackMsg := fmt.Sprintf("工具调用失败: %v\n\n我将基于已有知识尝试回答。", err)
			response, _ := a.SimpleChat(userMessage)
			return fallbackMsg + "\n\n" + response, nil
		}

		// 3. 基于工具结果生成最终回答
		prompt := fmt.Sprintf(`基于以下搜索结果，回答用户的问题。

搜索结果:
%s

用户问题: %s

请整合以上信息，给出清晰、准确的回答。如果搜索结果中有冲突或不确定的信息，请明确指出。`,
			toolResult, userMessage)

		return a.SimpleChat(prompt)
	}

	// 4. 不需要工具，直接回答
	return a.SimpleChat(userMessage)
}

// analyzeIntent 分析用户意图，判断是否需要调用工具
func (a *AgentService) analyzeIntent(message string) (bool, string, map[string]interface{}) {
	// 检查是否有游戏攻略相关的工具
	gameGuideTool, hasGameTool := a.toolRegistry.GetTool("game_guide_search")
	if !hasGameTool {
		return false, "", nil
	}

	// 使用游戏攻略工具的判断逻辑
	if guideTool, ok := gameGuideTool.(*tools.GameGuideTool); ok {
		if guideTool.IsGameRelated(message) {
			// 提取游戏信息
			gameName, topic := guideTool.ExtractGameInfo(message)

			args := map[string]interface{}{
				"game_name": gameName,
				"topic":     topic,
			}

			return true, "game_guide_search", args
		}
	}

	return false, "", nil
}
