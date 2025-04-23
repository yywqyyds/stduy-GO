package router

import (
	"homework/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//路由注册
	r.POST("/api/questions/create", handler.GenerateHandler)

	return r
}
