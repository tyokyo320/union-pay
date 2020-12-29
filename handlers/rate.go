package handlers

import (
	"fmt"
	"net/http"
	"union-pay/global"
	"union-pay/repository"

	"github.com/gin-gonic/gin"
)

type RateRequest struct {
	Date     string `json:"date"`
	Page     int    `json:page`
	pageSize int    `json:pageSize`
}

// ShowIndexPage: 用于展示主页
func ShowIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "hello world!",
		},
	)
}

// GetRate: 获取具体某一天的汇率
// handlers -> repository
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

// GetHistoryRate获取最近n天的历史汇率
func GetHistoryRate(c *gin.Context) {
	var form RateRequest
	err := c.ShouldBindJSON(&form)
	if err != nil {
		fmt.Println("form err", err)
	}
	fmt.Println(form.Page)
	fmt.Println(form.pageSize)

	var repo *repository.RateRepository = repository.NewRateRepository(global.POSTGRESQL_DB)
	historyRate, err := repo.ReadList(form.Page, form.pageSize)

	if err != nil {
		fmt.Println("Unable to get the historical exchange rate that was queried")
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"historyRate": historyRate,
			},
		)
	}

}
