package handler

import (
	"ai-agent/internal/agent"
	"ai-agent/internal/config"
	"ai-agent/pkg/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatRequest struct {
	Message  string              `json:"message,omitempty"`
	Messages []agent.ChatMessage `json:"messages,omitempty"`
}

type ChatResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func SetupRouter(agentService *agent.AgentService, serverCfg config.ServerConfig) *gin.Engine {
	router := gin.Default()

	// 提供静态文件（前端页面）
	router.Static("/web", "./web")
	router.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})
	router.GET("/search", func(c *gin.Context) {
		c.File("./web/simple-search.html")
	})
	router.GET("/smart", func(c *gin.Context) {
		c.File("./web/smart-search.html")
	})
	router.GET("/agent", func(c *gin.Context) {
		c.File("./web/full-agent.html")
	})

	// 设置超时
	router.Use(func(c *gin.Context) {
		timeout := time.Duration(serverCfg.WriteTimeout) * time.Second
		c.Set("timeout", timeout)
		c.Next()
	})

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// API路由组
	api := router.Group("/api/v1")
	{
		api.POST("/chat", handleChat(agentService))
		api.POST("/simple-chat", handleSimpleChat(agentService))
		api.POST("/smart-chat", handleSmartChat(agentService)) // ⭐ 新增：智能对话
		api.GET("/tools", handleListTools(agentService))       // ⭐ 新增：列出工具
	}

	return router
}

func handleChat(agentService *agent.AgentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ChatRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, ChatResponse{
				Success: false,
				Error:   "Invalid request: " + err.Error(),
			})
			return
		}

		var messages []agent.ChatMessage
		
		// 如果提供了 messages 数组，使用它（多轮对话）
		if len(req.Messages) > 0 {
			// 验证消息数量限制
			if len(req.Messages) > 20 {
				c.JSON(http.StatusBadRequest, ChatResponse{
					Success: false,
					Error:   "Too many messages, maximum is 20",
				})
				return
			}
			
			messages = req.Messages
			logger.Info("Received multi-turn chat request with " + fmt.Sprintf("%d", len(messages)) + " messages")
		} else if req.Message != "" {
			// 验证消息长度
			if len(req.Message) > 1000 {
				c.JSON(http.StatusBadRequest, ChatResponse{
					Success: false,
					Error:   "Message too long, maximum is 1000 characters",
				})
				return
			}
			
			// 否则使用单条消息（向后兼容）
			messages = []agent.ChatMessage{
				{Role: "user", Content: req.Message},
			}
			logger.Info("Received single chat request: " + req.Message)
		} else {
			c.JSON(http.StatusBadRequest, ChatResponse{
				Success: false,
				Error:   "Either 'message' or 'messages' must be provided",
			})
			return
		}

		response, err := agentService.Chat(messages)
		if err != nil {
			logger.Error("Chat error: " + err.Error())
			c.JSON(http.StatusInternalServerError, ChatResponse{
				Success: false,
				Error:   "Failed to get response: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, ChatResponse{
			Success: true,
			Data:    response,
		})
	}
}

func handleSimpleChat(agentService *agent.AgentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ChatRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, ChatResponse{
				Success: false,
				Error:   "Invalid request: " + err.Error(),
			})
			return
		}

		logger.Info("Received simple chat request: " + req.Message)

		response, err := agentService.SimpleChat(req.Message)
		if err != nil {
			logger.Error("Simple chat error: " + err.Error())
			c.JSON(http.StatusInternalServerError, ChatResponse{
				Success: false,
				Error:   "Failed to get response: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, ChatResponse{
			Success: true,
			Data:    response,
		})
	}
}

// handleSmartChat 智能对话处理器（支持工具调用）
func handleSmartChat(agentService *agent.AgentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ChatRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, ChatResponse{
				Success: false,
				Error:   "Invalid request: " + err.Error(),
			})
			return
		}

		logger.Info("Received smart chat request: " + req.Message)

		response, err := agentService.SmartChat(req.Message)
		if err != nil {
			logger.Error("Smart chat error: " + err.Error())
			c.JSON(http.StatusInternalServerError, ChatResponse{
				Success: false,
				Error:   "Failed to get response: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, ChatResponse{
			Success: true,
			Data:    response,
		})
	}
}

// handleListTools 列出可用工具
func handleListTools(agentService *agent.AgentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 需要从 agentService 获取工具列表
		// 暂时返回空列表
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []interface{}{},
		})
	}
}
