package router

import (
	"homework/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//调用大模型出题
	r.POST("/api/questions/create", handler.GenerateHandler)

	return r
}
