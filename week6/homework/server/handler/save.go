package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"server/db"
	"server/schema"

	"github.com/gin-gonic/gin"
)

func SaveQuestionsHandler(c *gin.Context) {
	var questions []schema.Question

	// 尝试绑定为数组
	if err := c.ShouldBindJSON(&questions); err != nil {
		// 如果不是数组，尝试绑定为单个对象
		var single schema.Question
		if err := c.ShouldBindJSON(&single); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
		// 包装成数组
		questions = append(questions, single)
	}

	// 开始事务
	tx, err := db.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to begin transaction"})
		return
	}

	stmt, err := tx.Prepare(`
		INSERT INTO questions (
			model, language, type, keyword, difficulty,
			question, options, answer, explanation,
			ai_start_time, ai_end_time, ai_cost_time
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to prepare statement"})
		return
	}
	defer stmt.Close()

	for _, q := range questions {
		optionsJSON, _ := json.Marshal(q.AiRes.Options)
		answerJSON, _ := json.Marshal(q.AiRes.Answer)

		_, err := stmt.Exec(
			q.AiReq.Model, q.AiReq.Language, q.AiReq.Type, q.AiReq.Keyword, q.AiReq.Difficulty,
			q.AiRes.Title, string(optionsJSON), string(answerJSON), q.AiRes.Explanation,
			q.AiStartTime, q.AiEndTime, q.AiCostTime,
		)
		if err != nil {
			log.Println("保存题目失败:", err)
			continue
		}
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "questions saved successfully"})
}
