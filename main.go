package main

import (
	"union-pay/config"
	"union-pay/global"
	"union-pay/initialize"
	"union-pay/routes"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	// Singleton pattern: 给全局变量赋值（给实例化的变量赋值）
	// 先读取配置文件
	global.CONFIG = config.NewConfig(".")
	// 然后才能初始化连接数据库
	global.POSTGRESQL_DB = initialize.NewGorm(global.CONFIG.PostGreSQL)
}

func main() {

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	routes.InitializeRoutes(router)

	router.Run(":8080")

}
