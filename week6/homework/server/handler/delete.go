package handler

import (
	"net/http"
	"server/db"

	"github.com/gin-gonic/gin"
)

type DeleteRequest struct {
	IDs []int `json:"ids"`
}

func DeleteQuestionsHandler(c *gin.Context) {
	var req DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.IDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误，缺少ids",
		})
		return
	}

	tx, err := db.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -2,
			"msg":  "开启事务失败",
		})
		return
	}

	stmt, err := tx.Prepare("DELETE FROM questions WHERE id = ?")
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -3,
			"msg":  "预处理删除语句失败",
		})
		return
	}
	defer stmt.Close()

	var notFoundIDs []int
	var deletedCount int

	for _, id := range req.IDs {
		res, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": -4,
				"msg":  "删除执行失败",
			})
			return
		}
		affected, _ := res.RowsAffected()
		if affected == 0 {
			notFoundIDs = append(notFoundIDs, id)
		} else {
			deletedCount++
		}
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -5,
			"msg":  "事务提交失败",
		})
		return
	}

	if deletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -6,
			"msg":  "未找到任何匹配的题目",
			"data": gin.H{"not_found_ids": notFoundIDs},
		})
		return
	}

	msg := "删除成功"
	if len(notFoundIDs) > 0 {
		msg = "部分题目未找到，其余已删除"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"data": gin.H{
			"deleted":       deletedCount,
			"not_found_ids": notFoundIDs,
		},
	})
}
