package handler

import (
	"encoding/json"
	"net/http"
	"server/db"
	"server/schema"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ModifyQuestionHandler handles POST /api/questions/modify/:id
func ModifyQuestionsHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 ID"})
		return
	}

	var q schema.Question
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求体不合法"})
		return
	}

	// 编码 options 和 answer 为 JSON 字符串
	optionsJSON, _ := json.Marshal(q.AiRes.Options)
	answerJSON, _ := json.Marshal(q.AiRes.Answer)

	res, err := db.DB.Exec(`
		UPDATE questions SET 
			model = ?, language = ?, type = ?, keyword = ?,
			question = ?, options = ?, answer = ?, explanation = ?,
			ai_start_time = ?, ai_end_time = ?, ai_cost_time = ?
		WHERE id = ?
	`, q.AiReq.Model, q.AiReq.Language, q.AiReq.Type, q.AiReq.Keyword,
		q.AiRes.Title, string(optionsJSON), string(answerJSON), q.AiRes.Explanation,
		q.AiStartTime, q.AiEndTime, q.AiCostTime, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库更新失败"})
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "题目未找到"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "题目修改成功"})
}
