package routes

import (
	"union-pay/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/", handlers.ShowIndexPage)
	rateRoutes := router.Group("/rate")
	{
		rateRoutes.POST("/read", handlers.GetRate)
		rateRoutes.POST("/latest", handlers.GetLatestRate)
		rateRoutes.POST("/history", handlers.GetHistoryRate)
	}
}
