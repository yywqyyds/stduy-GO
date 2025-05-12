package schema

type AiRequest struct {
	Model      string `json:"model"`      // 指定使用模型（可选）
	Language   string `json:"language"`   // 编程语言（必选）
	Type       int    `json:"type"`       // 题型：1 单选，2 多选
	Keyword    string `json:"keyword"`    // 提示关键词（必选）
	Count      int    `json:"count"`      // 题目数量
	Difficulty string `json:"difficulty"` // 题目难度
}

type AiRespond struct {
	Title       string   `json:"question"`    // 题干
	Options     []string `json:"options"`     // 四个选项
	Answer      []string `json:"answer"`      // 正确答案（支持多选）
	Explanation string   `json:"explanation"` // 解析
}

type Question struct {
	AiStartTime string    `json:"aiStartTime"` // 创建时间
	AiEndTime   string    `json:"aiEndTime"`   //结束时间
	AiCostTime  int       `json:"aiCostTime"`  //花费时间
	AiReq       AiRequest `json:"aiReq"`       //用户请求参数
	AiRes       AiRespond `json:"aiRes"`       //大模型返回JSON
}
