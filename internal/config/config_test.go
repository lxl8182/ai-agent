package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// 创建临时配置文件
	content := `
server:
  host: "localhost"
  port: 8080
agent:
  model: "gpt-3.5-turbo"
  api_key: "test-key"
  base_url: "https://api.openai.com/v1"
log:
  level: "info"
  file: "test.log"
`
	
	tmpfile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// 测试加载配置
	cfg, err := LoadConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// 验证配置值
	if cfg.Server.Host != "localhost" {
		t.Errorf("Expected host localhost, got %s", cfg.Server.Host)
	}
	if cfg.Server.Port != 8080 {
		t.Errorf("Expected port 8080, got %d", cfg.Server.Port)
	}
	if cfg.Agent.Model != "gpt-3.5-turbo" {
		t.Errorf("Expected model gpt-3.5-turbo, got %s", cfg.Agent.Model)
	}
}

func TestLoadConfigInvalidFile(t *testing.T) {
	_, err := LoadConfig("nonexistent-file.yaml")
	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
}
