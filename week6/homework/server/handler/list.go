package handler

import (
	"net/http"
	"server/db"
	"strings"

	"github.com/gin-gonic/gin"
)

type SimpleQuestion struct {
	Question   string `json:"question"`
	Type       int    `json:"type"`
	Difficulty string `json:"difficulty"`
}

// GetQuestionsHandler 获取题目列表（仅返回题干、题型、难度）
func GetQuestionsHandler(c *gin.Context) {
	typeParam := c.Query("type") // 题型筛选
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")

	query := `SELECT question, type, difficulty FROM questions`
	conditions := []string{}
	args := []interface{}{}

	if typeParam != "" {
		conditions = append(conditions, "type = ?")
		args = append(args, typeParam)
	}
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += " ORDER BY id DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": "数据库查询失败"})
		return
	}
	defer rows.Close()

	var list []SimpleQuestion
	for rows.Next() {
		var q SimpleQuestion
		if err := rows.Scan(&q.Question, &q.Type, &q.Difficulty); err == nil {
			list = append(list, q)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": list,
	})
}
