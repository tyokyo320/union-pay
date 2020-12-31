// 运行整个项目
package main

import (
	"union-pay/config"
	"union-pay/global"
	"union-pay/initialize"
	"union-pay/routers"
	"union-pay/tasks"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	// Singleton pattern: 给全局变量赋值（给实例化的变量赋值）
	// 先读取配置文件
	global.CONFIG = config.NewConfig(".")
	// 然后才能初始化连接数据库
	global.POSTGRESQL_DB = initialize.NewGorm(global.CONFIG.PostGreSQL)
	// Redis
	global.REDIS = initialize.NewRedis(global.CONFIG.Redis)
}

func main() {

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	routers.InitializeRouters(router)
	routers.InitializeRateRouters(router)

	tasks.RunTasks()

	router.Run(":8080")

}
