package handlers

import (
	"fmt"
	"net/http"
	"union-pay/global"
	"union-pay/repository"

	"github.com/gin-gonic/gin"
)

type RateRequest struct {
	Date string `json:"date"`
}

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
// 获取具体某一天的汇率
func GetRate(c *gin.Context) {
	var form RateRequest
	err := c.ShouldBindJSON(&form)
	if err != nil {
		fmt.Println("form err", err)
	}
	fmt.Println(form.Date)

	var repo *repository.RateRepository = repository.NewRateRepository(global.POSTGRESQL_DB)
	rate := repo.Read(form.Date)

	if rate == nil {
		fmt.Println("Current exchange rate query display to be updated")
	} else {
		// c.JSON(http.StatusOK, rate)

		c.JSON(
			http.StatusOK,
			gin.H{
				"rate": rate,
			},
		)
	}

}

// GetLatestRate获取最近一天的汇率
func GetLatestRate(c *gin.Context) {
	var repo *repository.RateRepository = repository.NewRateRepository(global.POSTGRESQL_DB)
	lastestRate := repo.ReadLastest()

	c.JSON(
		http.StatusOK,
		gin.H{
			"lastestRate": lastestRate,
		},
	)

}

func GetHistoryRate(c *gin.Context) {
	var repo *repository.RateRepository = repository.NewRateRepository(global.POSTGRESQL_DB)

	day := 3
	historyRate := repo.ReadHistory(day)

	c.JSON(
		http.StatusOK,
		gin.H{
			"historyRate": historyRate,
		},
	)
}
