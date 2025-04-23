package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/sum", func(c *gin.Context) {
		fmt.Print(c)
		// 获取 query 参数
		xStr := c.Query("x")
		yStr := c.Query("y")

		// 转换为整数
		x, err1 := strconv.Atoi(xStr)
		y, err2 := strconv.Atoi(yStr)

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "参数 x 和 y 必须是整数",
			})
			return
		}

		// 返回 JSON
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"reqData": gin.H{
				"x": x,
				"y": y,
			},
			"data": x + y,
		})
	})

	r.Run(":8080") // 启动服务
}
