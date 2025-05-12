// handler/generate.go
package handler

import (
	"net/http"
	"server/service"
	"strings"

	"github.com/gin-gonic/gin"
)

// 请求参数结构体
type req struct {
	Model      string `json:"model"`      // 指定使用模型（可选）
	Language   string `json:"language"`   // 编程语言（必选）
	Type       int    `json:"type"`       // 题型：1 单选，2 多选
	Keyword    string `json:"keyword"`    // 提示关键词（必选）
	Difficulty string `json:"difficulty"` // 难度（简单、中等、困难）
	Count      int    `json:"count"`      // 题目数量
}

// 题型转换映射表
var typeMap = map[int]string{
	1: "单选题",
	2: "多选题",
	3: "编程题",
}

// 模型转换映射表
var typeModel = map[string]string{
	"tongyi":   "qwen-max",
	"deepseek": "deepseek-v3",
}

// 支持的语言种类，不是这五种将默认设为go
var supportedLanguages = map[string]bool{
	"go":         true,
	"javascript": true,
	"java":       true,
	"python":     true,
	"c++":        true,
}

// GenerateHandler 处理 /generate 请求
func GenerateHandler(c *gin.Context) {
	var req req
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": "请求参数错误"})
		return
	}

	language := strings.ToLower(req.Language)
	if !supportedLanguages[language] {
		language = "go" // 默认语言
	}

	qType, ok := typeMap[req.Type]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": -3, "msg": "不支持的题型类型"})
		return
	}

	// 设置默认 model 和 language
	if strings.TrimSpace(req.Model) == "" {
		req.Model = "tongyi"
	}
	qModel, ok := typeModel[req.Model]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": -4, "msg": "不支持的大模型类型"})
		return
	}

	questions, err := service.CallLLM(qModel, req.Language, req.Keyword, req.Difficulty, qType, req.Count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -5, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "生成成功",
		"data": questions,
	})
}
