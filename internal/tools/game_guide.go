package tools

import (
	"fmt"
	"strings"
)

// GameGuideTool 游戏攻略搜索工具（专门优化）
type GameGuideTool struct {
	searchTool *WebSearchTool
}

// NewGameGuideTool 创建游戏攻略工具
func NewGameGuideTool(searchTool *WebSearchTool) *GameGuideTool {
	return &GameGuideTool{
		searchTool: searchTool,
	}
}

func (t *GameGuideTool) Name() string {
	return "game_guide_search"
}

func (t *GameGuideTool) Description() string {
	return "专门用于搜索游戏攻略、 walkthrough、技巧、boss打法等游戏相关信息"
}

func (t *GameGuideTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"game_name": map[string]interface{}{
			"type":        "string",
			"description": "游戏名称，例如: '原神'、'黑神话悟空'",
			"required":    true,
		},
		"topic": map[string]interface{}{
			"type":        "string",
			"description": "攻略主题，例如: '角色培养'、'boss打法'、'隐藏要素'",
			"required":    true,
		},
		"platform": map[string]interface{}{
			"type":        "string",
			"description": "游戏平台（可选），例如: 'PC'、'PS5'、'Switch'",
			"required":    false,
		},
	}
}

// Execute 执行游戏攻略搜索
func (t *GameGuideTool) Execute(args map[string]interface{}) (string, error) {
	gameName, ok := args["game_name"].(string)
	if !ok || gameName == "" {
		return "", fmt.Errorf("game_name is required")
	}

	topic, ok := args["topic"].(string)
	if !ok || topic == "" {
		return "", fmt.Errorf("topic is required")
	}

	platform, _ := args["platform"].(string)

	// 构建优化的搜索查询
	query := t.buildGameGuideQuery(gameName, topic, platform)

	// 调用通用搜索工具
	result, err := t.searchTool.Execute(map[string]interface{}{
		"query":       query,
		"num_results": 8, // 游戏攻略需要更多结果
	})

	if err != nil {
		return "", err
	}

	// 添加额外的提示信息
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("🎮 **%s - %s** 攻略搜索结果:\n\n", gameName, topic))
	sb.WriteString(result)
	sb.WriteString("\n💡 **提示**: 以上是从网络搜索的最新攻略信息，建议结合多个来源获取最准确的信息。\n")

	return sb.String(), nil
}

// buildGameGuideQuery 构建优化的游戏攻略搜索查询
func (t *GameGuideTool) buildGameGuideQuery(gameName, topic, platform string) string {
	var parts []string

	// 基础查询：游戏名 + 攻略
	parts = append(parts, gameName)
	parts = append(parts, "攻略")

	// 添加主题
	if topic != "" {
		parts = append(parts, topic)
	}

	// 添加平台
	if platform != "" {
		parts = append(parts, platform)
	}

	// 添加一些常用的攻略关键词
	guideKeywords := []string{
		"详细攻略",
		"完整指南",
		"新手教程",
		"技巧",
	}

	// 随机选择一个关键词（这里简化为固定使用第一个）
	parts = append(parts, guideKeywords[0])

	query := strings.Join(parts, " ")

	return query
}

// IsGameRelated 判断问题是否与游戏相关
func (t *GameGuideTool) IsGameRelated(message string) bool {
	// 游戏相关的关键词
	gameKeywords := []string{
		"攻略", "walkthrough", "guide",
		"怎么玩", "怎么过", "怎么打",
		"boss", "副本", "关卡",
		"角色", "装备", "技能",
		"任务", "剧情", "结局",
		"隐藏", "彩蛋", "成就",
		"升级", "培养", "build",
	}

	messageLower := strings.ToLower(message)

	for _, keyword := range gameKeywords {
		if strings.Contains(messageLower, strings.ToLower(keyword)) {
			return true
		}
	}

	return false
}

// ExtractGameInfo 从用户消息中提取游戏信息
func (t *GameGuideTool) ExtractGameInfo(message string) (gameName string, topic string) {
	// 这是一个简化的实现
	// 在实际生产中，可以使用 LLM 来更准确地提取

	// 常见的游戏名称模式
	// 例如: "原神 雷电将军怎么培养" -> 游戏: 原神, 主题: 雷电将军培养

	parts := strings.Fields(message)

	if len(parts) >= 2 {
		// 假设第一个词是游戏名
		gameName = parts[0]

		// 剩余部分是主题
		topic = strings.Join(parts[1:], " ")
	} else {
		// 如果只有一个词，可能是游戏名
		gameName = message
		topic = "攻略"
	}

	return gameName, topic
}
