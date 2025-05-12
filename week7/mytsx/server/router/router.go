package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//路由注册
	r.POST("/api/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "true",
			"data": nil,
		})
	})

	return r
}
