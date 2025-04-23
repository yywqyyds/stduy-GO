// handler/generate.go
package handler

import (
	"encoding/json"
	"homework/service"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 请求参数结构体
type GenerateRequest struct {
	Model    string `json:"model"`    // 指定使用模型（可选）
	Language string `json:"language"` // 编程语言（必选）
	Type     int    `json:"type"`     // 题型：1 单选，2 多选
	Keyword  string `json:"keyword"`  // 提示关键词（必选）
}

// 题型转换映射表
var typeMap = map[int]string{
	1: "单选题",
	2: "多选题",
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
	var req GenerateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  -1,
			"msg":   "请求参数错误",
			"aiRes": nil,
		})
		return
	}

	language := strings.ToLower(req.Language)
	if !supportedLanguages[language] {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  -2,
			"msg":   "不支持的语言类型",
			"aiRes": nil,
		})
		return
	}

	qType, ok := typeMap[req.Type]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  -3,
			"msg":   "不支持的题型类型",
			"aiRes": nil,
		})
		return
	}

	qModel, ok := typeModel[req.Model]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  -4,
			"msg":   "不支持的大模型类型",
			"aiRes": nil,
		})
		return
	}

	result, err := service.CallLLM(qModel, req.Language, req.Keyword, qType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -5,
			"msg":   err.Error(),
			"aiRes": nil,
		})
		return
	}

	//存储JSON文件到 data/目录下
	today := time.Now().Format("2006-01-02")
	filePath := filepath.Join("data", today+".json")
	// 确保 data 目录存在
	_ = os.MkdirAll("data", 0755)

	// 打开文件，准备追加写入
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -6,
			"msg":   "写入文件失败",
			"aiRes": nil,
		})
		return
	}
	defer f.Close()

	// 写入一条数据（追加）
	encoded, _ := json.Marshal(result)
	f.Write(append(encoded, '\n'))
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "",
		"aiRes": result,
	})
}
