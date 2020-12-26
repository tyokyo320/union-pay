package handlers

import (
	"net/http"
	"union-pay/global"
	"union-pay/repository"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "hello world!",
		},
	)
}

// handlers -> repository
func GetRate(c *gin.Context) {
	date := "2020-12-24"

	var repo *repository.RateRepository = repository.NewRateRepository(global.POSTGRESQL_DB)
	rate := repo.Read(date)

	c.JSON(http.StatusOK, rate)
}
