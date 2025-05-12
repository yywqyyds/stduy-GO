package main

import (
	"server/config"
	"server/db"
	"server/router"
)

func main() {
	//加载环境变量
	config.LoadConfig(".env")

	//初始化数据库
	db.InitDB()

	//初始化路由并运行
	r := router.SetupRouter()
	r.Run(":8080")

}
