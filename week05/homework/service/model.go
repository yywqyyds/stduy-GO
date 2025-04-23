// model/model.go
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"homework/config"
	"homework/schema"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// CallLLM 调用大模型生成题目
func CallLLM(modelOverride, lang, keyword, Type string) (schema.Question, error) {
	modelName := config.ModelName
	if modelOverride != "" {
		modelName = modelOverride
	}

	// 构造提示词
	prompt := fmt.Sprintf(`
			请生成一道%s类型的选择题，语言为 %s，关键词为 %s。只能输出 JSON 格式，禁止添加 markdown 或说明性文字。字段包括：
			question（题干）、options（选项 A/B/C/D）array类型、answer（正确答案）不管是单选题还是多选题都是array类型只包含选项、explanation（解析），必须按纯 JSON 返回。
			`, Type, lang, keyword)

	client := openai.NewClient(
		option.WithAPIKey(config.ApiKey),
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
	)
	aiStartTime := time.Now()
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(prompt),
			},
			Model: modelName,
		},
	)

	if err != nil {
		panic(err.Error())
	}

	responseContent := chatCompletion.Choices[0].Message.Content
	aiEndTime := time.Now()
	aiCostTime := aiEndTime.Sub(aiStartTime)
	fmt.Println(responseContent)

	// 解码内容为结构体
	var q schema.Question
	// 补全元信息

	var typeMap = map[string]int{
		"单选题": 1,
		"多选题": 2,
	}
	q.AiStartTime = aiStartTime.Format(time.RFC3339)
	q.AiEndTime = aiEndTime.Format(time.RFC3339)
	q.AiCostTime = int(aiCostTime.Seconds())
	q.AiReq = schema.AiRequest{
		Model:    modelName,
		Language: lang,
		Type:     typeMap[Type],
		Keyword:  keyword,
	}
	err = json.Unmarshal([]byte(responseContent), &q.AiRes)
	if err != nil {
		return schema.Question{}, fmt.Errorf("模型返回内容格式非法: %v", err)
	}

	fmt.Println(q)
	return q, nil
}
