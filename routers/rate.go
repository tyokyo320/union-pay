package routers

import (
	"union-pay/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRateRouters(router *gin.Engine) {
	rateRoutes := router.Group("/rate")
	{
		rateRoutes.POST("/read", handlers.GetRate)
		rateRoutes.POST("/latest", handlers.GetLatestRate)
		rateRoutes.POST("/history", handlers.GetHistoryRate)
	}
}
