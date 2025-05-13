package handler

import (
	"fmt"
	"net/http"
	"server/db"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type SimpleQuestion struct {
	ID         int    `json:"id"`
	Question   string `json:"question"`
	Type       int    `json:"type"`
	Difficulty string `json:"difficulty"`
}

func GetQuestionsHandler(c *gin.Context) {
	typeParam := c.Query("type")
	keyword := c.Query("keyword")
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	// 转换 limit 和 offset 为 int
	limit, err1 := strconv.Atoi(limitStr)
	offset, err2 := strconv.Atoi(offsetStr)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "分页参数格式错误",
		})
		return
	}

	query := `SELECT id, question, type, difficulty FROM questions`
	conditions := []string{}
	args := []interface{}{}

	// 类型筛选
	if typeParam != "" {
		typeInt, err := strconv.Atoi(typeParam)
		if err == nil {
			conditions = append(conditions, "type = ?")
			args = append(args, typeInt)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  "type 参数必须为数字",
			})
			return
		}
	}

	// 关键词筛选
	if keyword != "" {
		conditions = append(conditions, "question LIKE ?")
		args = append(args, "%"+keyword+"%")
	}

	// 拼接 WHERE
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY id DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	// 打印调试信息
	fmt.Println("【DEBUG】SQL:", query)
	fmt.Println("【DEBUG】ARGS:", args)

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		fmt.Println("【ERROR】查询失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "数据库查询失败",
		})
		return
	}
	defer rows.Close()

	var list []SimpleQuestion
	for rows.Next() {
		var q SimpleQuestion
		if err := rows.Scan(&q.ID, &q.Question, &q.Type, &q.Difficulty); err == nil {
			list = append(list, q)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": list,
	})
}
