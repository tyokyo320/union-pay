package routers

import (
	"union-pay/handlers"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

func InitializeRouters(router *gin.Engine) {
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/jobrunner/json", func(c *gin.Context) {
		c.JSON(200, jobrunner.StatusJson())
	})
	router.GET("/jobrunner/html", func(c *gin.Context) {
		c.HTML(200, "", jobrunner.StatusPage())

	})
}
