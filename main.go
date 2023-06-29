package main

import (
	"ET-order-mini-program/configs"
	"ET-order-mini-program/database"
	"ET-order-mini-program/internal/app"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: 加载配置文件
	env := gin.Mode()
	config, err := configs.LoadConfig(env)
	if err != nil {
		panic(err.Error())
	}

	// TODO: 初始化数据库
	db := database.InitDB(config)

	// TODO: 初始化路由
	router := app.InitRouter(db)

	// TODO: 启动服务器
	app.Serve(router, config)
}
