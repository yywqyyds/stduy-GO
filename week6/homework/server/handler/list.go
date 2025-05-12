package handler

import (
	"net/http"
	"server/db"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetQuestionsHandler 获取题目列表接口
func GetQuestionsHandler(c *gin.Context) {
	typeParam := c.Query("type")  // 题型筛选（1=单选，2=多选，3=编程）
	keyword := c.Query("keyword") // 题干关键词模糊搜索
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")

	// 构建查询 SQL
	baseSQL := `SELECT id, model, language, type, keyword, question, options, answer, explanation, ai_start_time, ai_end_time, ai_cost_time FROM questions`
	where := []string{}
	args := []interface{}{}

	if typeParam != "" {
		where = append(where, "type = ?")
		args = append(args, typeParam)
	}
	if keyword != "" {
		where = append(where, "question LIKE ?")
		args = append(args, "%"+keyword+"%")
	}

	if len(where) > 0 {
		baseSQL += " WHERE " + strings.Join(where, " AND ")
	}
	baseSQL += " ORDER BY id DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := db.DB.Query(baseSQL, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": "数据库查询失败", "data": nil})
		return
	}
	defer rows.Close()

	// 构造响应结构
	type Question struct {
		ID          int    `json:"id"`
		Model       string `json:"model"`
		Language    string `json:"language"`
		Type        int    `json:"type"`
		Keyword     string `json:"keyword"`
		Question    string `json:"question"`
		Options     string `json:"options"`
		Answer      string `json:"answer"`
		Explanation string `json:"explanation"`
		StartTime   string `json:"ai_start_time"`
		EndTime     string `json:"ai_end_time"`
		CostTime    int    `json:"ai_cost_time"`
	}

	var questions []Question
	for rows.Next() {
		var q Question
		err := rows.Scan(&q.ID, &q.Model, &q.Language, &q.Type, &q.Keyword, &q.Question, &q.Options, &q.Answer, &q.Explanation, &q.StartTime, &q.EndTime, &q.CostTime)
		if err == nil {
			questions = append(questions, q)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": questions,
	})
}
