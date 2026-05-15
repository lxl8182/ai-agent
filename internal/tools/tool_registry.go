package tools

import (
	"fmt"
	"sync"
)

// Tool 工具接口
type Tool interface {
	Name() string
	Description() string
	Parameters() map[string]interface{}
	Execute(args map[string]interface{}) (string, error)
}

// Registry 工具注册中心
type Registry struct {
	tools map[string]Tool
	mu    sync.RWMutex
}

// NewRegistry 创建工具注册中心
func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]Tool),
	}
}

// Register 注册工具
func (r *Registry) Register(tool Tool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tools[tool.Name()] = tool
}

// Execute 执行工具
func (r *Registry) Execute(name string, args map[string]interface{}) (string, error) {
	r.mu.RLock()
	tool, exists := r.tools[name]
	r.mu.RUnlock()

	if !exists {
		return "", fmt.Errorf("tool %s not found", name)
	}

	return tool.Execute(args)
}

// ListTools 列出所有可用工具
func (r *Registry) ListTools() []map[string]interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []map[string]interface{}
	for _, tool := range r.tools {
		result = append(result, map[string]interface{}{
			"name":        tool.Name(),
			"description": tool.Description(),
			"parameters":  tool.Parameters(),
		})
	}

	return result
}

// GetTool 获取指定工具
func (r *Registry) GetTool(name string) (Tool, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tool, exists := r.tools[name]
	return tool, exists
}
