package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// WebSearchTool 网络搜索工具
type WebSearchTool struct {
	apiKey string
	engine string // "bing" or "google"
}

// SearchResult 搜索结果
type SearchResult struct {
	Title       string `json:"title"`
	Snippet     string `json:"snippet"`
	Link        string `json:"link"`
	PublishedAt string `json:"published_at,omitempty"`
}

// NewWebSearchTool 创建网络搜索工具
func NewWebSearchTool(apiKey string, engine string) *WebSearchTool {
	if engine == "" {
		engine = "bing" // 默认使用 Bing
	}
	return &WebSearchTool{
		apiKey: apiKey,
		engine: engine,
	}
}

func (t *WebSearchTool) Name() string {
	return "web_search"
}

func (t *WebSearchTool) Description() string {
	return "搜索互联网获取最新信息，特别适用于查询游戏攻略、新闻、实时数据等"
}

func (t *WebSearchTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"query": map[string]interface{}{
			"type":        "string",
			"description": "搜索关键词，例如: '原神 4.5版本攻略'",
			"required":    true,
		},
		"num_results": map[string]interface{}{
			"type":        "integer",
			"description": "返回结果数量，默认5",
			"default":     5,
		},
	}
}

// Execute 执行搜索
func (t *WebSearchTool) Execute(args map[string]interface{}) (string, error) {
	query, ok := args["query"].(string)
	if !ok || query == "" {
		return "", fmt.Errorf("search query is required")
	}

	numResults := 5
	if n, ok := args["num_results"].(float64); ok {
		numResults = int(n)
	}

	var results []SearchResult
	var err error

	// 根据搜索引擎选择调用不同的API
	switch t.engine {
	case "bing":
		results, err = t.searchWithBing(query, numResults)
	case "google":
		results, err = t.searchWithGoogle(query, numResults)
	default:
		results, err = t.searchWithBing(query, numResults)
	}

	if err != nil {
		return "", fmt.Errorf("search failed: %w", err)
	}

	if len(results) == 0 {
		return "未找到相关搜索结果", nil
	}

	// 格式化搜索结果
	return t.formatResults(results), nil
}

// searchWithBing 使用 Bing Search API
func (t *WebSearchTool) searchWithBing(query string, count int) ([]SearchResult, error) {
	// Bing Search API endpoint
	baseURL := "https://api.bing.microsoft.com/v7.0/search"

	// 构建请求
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 添加查询参数
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("count", fmt.Sprintf("%d", count))
	q.Add("mkt", "zh-CN") // 中文市场
	req.URL.RawQuery = q.Encode()

	// 添加认证头
	req.Header.Set("Ocp-Apim-Subscription-Key", t.apiKey)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Bing API returned status %d: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var bingResp BingSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&bingResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// 转换结果
	var results []SearchResult
	for _, item := range bingResp.WebPages.Value {
		results = append(results, SearchResult{
			Title:   item.Name,
			Snippet: item.Snippet,
			Link:    item.URL,
		})
	}

	return results, nil
}

// searchWithGoogle 使用 Google Custom Search API
func (t *WebSearchTool) searchWithGoogle(query string, count int) ([]SearchResult, error) {
	// Google Custom Search API endpoint
	baseURL := "https://www.googleapis.com/customsearch/v1"

	// 注意：这里需要额外的 CX (Custom Search Engine ID)
	// 为简化，暂时返回错误，提示配置
	if t.apiKey == "" {
		return nil, fmt.Errorf("Google Search requires both API key and CX ID")
	}

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	q := req.URL.Query()
	q.Add("key", t.apiKey)
	q.Add("cx", "YOUR_CX_ID") // 需要替换为实际的 CX ID
	q.Add("q", query)
	q.Add("num", fmt.Sprintf("%d", count))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var googleResp GoogleSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&googleResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var results []SearchResult
	for _, item := range googleResp.Items {
		results = append(results, SearchResult{
			Title:   item.Title,
			Snippet: item.Snippet,
			Link:    item.Link,
		})
	}

	return results, nil
}

// formatResults 格式化搜索结果为文本
func (t *WebSearchTool) formatResults(results []SearchResult) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("找到 %d 个相关结果:\n\n", len(results)))

	for i, result := range results {
		sb.WriteString(fmt.Sprintf("%d. **%s**\n", i+1, result.Title))
		sb.WriteString(fmt.Sprintf("   %s\n", result.Snippet))
		sb.WriteString(fmt.Sprintf("   链接: %s\n\n", result.Link))
	}

	return sb.String()
}

// BingSearchResponse Bing API 响应结构
type BingSearchResponse struct {
	WebPages struct {
		Value []struct {
			Name    string `json:"name"`
			Snippet string `json:"snippet"`
			URL     string `json:"url"`
		} `json:"value"`
	} `json:"webPages"`
}

// GoogleSearchResponse Google API 响应结构
type GoogleSearchResponse struct {
	Items []struct {
		Title   string `json:"title"`
		Snippet string `json:"snippet"`
		Link    string `json:"link"`
	} `json:"items"`
}
