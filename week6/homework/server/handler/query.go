package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/db"
	"server/schema"
	"strconv"

	"github.com/gin-gonic/gin"
)

// QueryQuestionHandler handles GET /api/questions/query/:id
func QueryQuestionsHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的题目 ID"})
		return
	}

	row := db.DB.QueryRow(`
		SELECT 
			model, language, type, keyword,
			question, options, answer, explanation,
			ai_start_time, ai_end_time, ai_cost_time
		FROM questions WHERE id = ?
	`, id)

	var q schema.Question
	var optionsStr, answerStr string

	err = row.Scan(
		&q.AiReq.Model, &q.AiReq.Language, &q.AiReq.Type, &q.AiReq.Keyword,
		&q.AiRes.Title, &optionsStr, &answerStr, &q.AiRes.Explanation,
		&q.AiStartTime, &q.AiEndTime, &q.AiCostTime,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "题目不存在"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询题目失败"})
		return
	}

	// 反序列化选项和答案字段
	if err := json.Unmarshal([]byte(optionsStr), &q.AiRes.Options); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "选项字段解析失败"})
		return
	}
	if err := json.Unmarshal([]byte(answerStr), &q.AiRes.Answer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "答案字段解析失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": q,
	})
}
