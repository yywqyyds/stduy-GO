package main

import (
	"log"

	"os"
	"user-manage/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 检查user.json文件是否存在
	if _, err := os.Stat("user.json"); os.IsNotExist(err) {
		log.Println("user.json文件不存在,创建文件")
		_, err := os.Create("user.json")
		if err != nil {
			log.Fatalf("创建文件失败: %v", err)
		}
	}

	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUser)
	router.POST("/users/update", handlers.UpdateUser)
	router.POST("/users/delete", handlers.DeleteUser)

	log.Println("服务器启动，监听端口 :8080")
	router.Run(":8080")
}
