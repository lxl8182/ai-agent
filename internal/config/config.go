package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Agent  AgentConfig  `yaml:"agent"`
	Log    LogConfig    `yaml:"log"`
}

type ServerConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type AgentConfig struct {
	Model        string  `yaml:"model"`
	APIKey       string  `yaml:"api_key"`
	BaseURL      string  `yaml:"base_url"`
	MaxTokens    int     `yaml:"max_tokens"`
	Temperature  float64 `yaml:"temperature"`
	SystemPrompt string  `yaml:"system_prompt"`
}

type LogConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
