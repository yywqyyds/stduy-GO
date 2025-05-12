package service

import (
	"context"
	"encoding/json"
	"fmt"
	"server/config"
	"server/schema"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// CallLLM 支持一次生成多个题目
func CallLLM(modelOverride, lang, keyword, difficulty, Type string, count int) ([]schema.Question, error) {
	modelName := config.ModelName
	if modelOverride != "" {
		modelName = modelOverride
	}

	var typeMap = map[string]int{
		"单选题": 1,
		"多选题": 2,
		"编程题": 3,
	}
	_, ok := typeMap[Type]
	if !ok {
		return nil, fmt.Errorf("不支持的题型: %s", Type)
	}

	// 构造提示词
	prompt := fmt.Sprintf(`
	请生成%d道%s类型的题目，语言为 %s，关键词为 %s，难度为 %s。只能输出 JSON 格式，禁止添加 markdown 或说明性文字。字段包括：
	question（题干）、options（选项 A/B/C/D）array类型、answer（正确答案）不管是单选题还是多选题都是array类型只包
	含选项、explanation（解析），如果是编程题只有题干，不需要答案。必须按纯 JSON 返回。
	`, count, Type, lang, keyword, difficulty)

	client := openai.NewClient(
		option.WithAPIKey(config.ApiKey),
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
	)

	start := time.Now()
	resp, err := client.Chat.Completions.New(
		context.TODO(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(prompt),
			},
			Model: modelName,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("调用大模型失败: %v", err)
	}
	end := time.Now()
	duration := end.Sub(start)

	// 打印原始响应内容
	content := resp.Choices[0].Message.Content
	fmt.Println("模型返回：", content)

	// 解码多个题目
	var questions []schema.Question

	if count == 1 {
		var singleResp schema.AiRespond
		if err := json.Unmarshal([]byte(content), &singleResp); err != nil {
			return nil, fmt.Errorf("模型返回内容格式非法（单题）: %v", err)
		}

		q := schema.Question{
			AiStartTime: start.Format(time.RFC3339),
			AiEndTime:   end.Format(time.RFC3339),
			AiCostTime:  int(duration.Milliseconds()),
			AiReq: schema.AiRequest{
				Model:    modelName,
				Language: lang,
				Type:     typeMap[Type],
				Keyword:  keyword,
			},
			AiRes: singleResp,
		}
		questions = append(questions, q)
	} else {
		var multiResp []schema.AiRespond
		if err := json.Unmarshal([]byte(content), &multiResp); err != nil {
			return nil, fmt.Errorf("模型返回内容格式非法（多题）: %v", err)
		}
		for _, item := range multiResp {
			q := schema.Question{
				AiStartTime: start.Format(time.RFC3339),
				AiEndTime:   end.Format(time.RFC3339),
				AiCostTime:  int(duration.Milliseconds()),
				AiReq: schema.AiRequest{
					Model:    modelName,
					Language: lang,
					Type:     typeMap[Type],
					Keyword:  keyword,
				},
				AiRes: item,
			}
			questions = append(questions, q)
		}
	}

	return questions, nil
}
