package router

import (
	"server/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	r.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	//调用大模型出题
	r.POST("/api/questions/create", handler.GenerateHandler)

	//获取题库列表
	r.GET("/api/questions/list", handler.GetQuestionsHandler)

	//查询题目
	r.GET("/api/questions/query/:id", handler.QueryQuestionsHandler)

	//删除题目
	r.POST("/api/questions/delete", handler.DeleteQuestionsHandler)

	//保存题目
	r.POST("/api/questions/save", handler.SaveQuestionsHandler)

	//编辑题目
	r.POST("/api/questions/modify/:id", handler.ModifyQuestionsHandler)

	return r
}
