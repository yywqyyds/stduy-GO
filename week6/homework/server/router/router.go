package router

import (
	"server/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//调用大模型出题
	r.POST("/api/questions/create", handler.GenerateHandler)

	//获取题库列表
	r.GET("/api/questions/list", handler.GetQuestionsHandler)

	//删除题目
	r.POST("/api/questions/delete", handler.DeleteQuestionsHandler)

	//保存题目
	r.POST("/api/questions/save", handler.SaveQuestionsHandler)

	return r
}
