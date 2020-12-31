package routers

import (
	"union-pay/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRouters(router *gin.Engine) {
	router.GET("/", handlers.ShowIndexPage)
}
