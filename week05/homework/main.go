package main

import (
	"homework/config"
	"homework/router"
)

func main(){
	//加载环境变量
	config.LoadConfig(".env")

	//初始化路由并运行
	r := router.SetupRouter()
	r.Run(":8080")
}