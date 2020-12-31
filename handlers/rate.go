package handlers

import (
	"fmt"
	"net/http"
	"union-pay/global"
	"union-pay/models"
	"union-pay/repository"

	"github.com/gin-gonic/gin"
)

// ShowIndexPage 用于展示主页
func ShowIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "hello world!",
		},
	)
}

// GetRate 获取具体某一天的汇率
// handlers -> repository
func GetRate(c *gin.Context) {
	var form models.RateRequest
	err := c.ShouldBindJSON(&form)
	if err != nil {
		fmt.Println("form err", err)
	}
	fmt.Println(form.Date)

	var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
	var redisRepo *repository.RateCacheRepository = repository.NewRateCacheRepository(global.REDIS)

	// 先从缓存中读取某一天汇率
	rate, err := redisRepo.Read(form.Date)
	if err != nil {
		// 如果缓存中没有，从Update DB中更新
		rateStruct := newRepo.Read(form.Date)
		// 然后添加至缓存中
		redisRepo.Create(form.Date, rateStruct.ExchangeRate)

		rate = rateStruct.ExchangeRate
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"rate": rate,
		},
	)
}

// GetLatestRate 获取最近一天的汇率
func GetLatestRate(c *gin.Context) {
	var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
	var redisRepo *repository.RateCacheRepository = repository.NewRateCacheRepository(global.REDIS)

	rate, err := redisRepo.Read("latest")
	if err != nil {
		lastestRate := newRepo.ReadLastest()
		redisRepo.Create("latest", lastestRate.ExchangeRate)
		rate = lastestRate.ExchangeRate
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"lastestRate": rate,
		},
	)
}

// GetHistoryRate 获取最近n天的历史汇率
func GetHistoryRate(c *gin.Context) {
	var form models.RateRequest
	err := c.BindJSON(&form)
	if err != nil {
		fmt.Println("form err", err)
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "invalid params",
			},
		)
		return
	}

	// fmt.Println(form.Page)
	// fmt.Println(form.PageSize)

	// cache list

	var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
	historyRate, err := newRepo.ReadList(form.Page, form.PageSize)

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
