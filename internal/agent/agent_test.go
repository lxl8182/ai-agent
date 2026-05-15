package agent

import (
	"ai-agent/internal/config"
	"testing"
)

func TestNewAgentService(t *testing.T) {
	cfg := config.AgentConfig{
		Model:        "gpt-3.5-turbo",
		APIKey:       "test-key",
		BaseURL:      "https://api.openai.com/v1",
		MaxTokens:    2000,
		Temperature:  0.7,
		SystemPrompt: "You are a helpful assistant.",
	}

	service := NewAgentService(cfg)
	if service == nil {
		t.Error("Expected agent service to be created")
	}
	if service.config.Model != cfg.Model {
		t.Errorf("Expected model %s, got %s", cfg.Model, service.config.Model)
	}
}

func TestSimpleChat(t *testing.T) {
	// 这是一个集成测试示例，需要有效的API密钥
	// 在实际使用中，应该使用mock进行测试
	t.Skip("Skipping integration test - requires valid API key")
}
