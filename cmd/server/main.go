// 运行整个项目
package main

import (
	"log"
	"os"
	"union-pay/config"
	"union-pay/global"
	"union-pay/initialize"
	"union-pay/routers"
	"union-pay/tasks"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	// Singleton pattern: 给全局变量赋值（给实例化的变量赋值）
	// 先读取配置文件
	global.CONFIG = config.NewConfig(".")
	// 然后才能初始化连接数据库
	global.POSTGRESQL_DB = initialize.NewGorm(global.CONFIG.PostGreSQL)
	// Redis
	global.REDIS = initialize.NewRedis(global.CONFIG.Redis)
	// Logger
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", 0700)
	}
	file, err := os.OpenFile("./logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	global.InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	global.WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	global.ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	global.PanicLogger = log.New(file, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	routers.InitializeRouters(router)
	routers.InitializeRateRouters(router)

	tasks.RunTasks()

	global.InfoLogger.Println("Starting the application...")
	router.Run(":8080")

}
