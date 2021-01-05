package handlers

import (
	"encoding/json"
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
		global.ErrorLogger.Println("[handlers]Form(GetRate) went wrong")
	}
	fmt.Println(form.Date)

	var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
	var redisRepo *repository.RateCacheRepository = repository.NewRateCacheRepository(global.REDIS)

	// 先从缓存中读取某一天汇率
	rateJSON, err := redisRepo.Read(form.Date)
	if err != nil {
		// 如果缓存中没有，从Update DB中更新
		rateStruct := newRepo.ReadLastest()
		j, err := json.Marshal(map[string]interface{}{
			"rate": rateStruct.ExchangeRate,
			"date": rateStruct.EffectiveDate,
		})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": "InternalServerError",
				},
			)
			global.ErrorLogger.Println("[handler]Json marshal went wrong")
			return
		}

		rateJSON = string(j)
		// 然后添加至缓存中
		redisRepo.Create(form.Date, string(j))
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteString(rateJSON)
}

// GetLatestRate 获取最近一天的汇率
func GetLatestRate(c *gin.Context) {
	var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
	var redisRepo *repository.RateCacheRepository = repository.NewRateCacheRepository(global.REDIS)

	rateJSON, err := redisRepo.Read("latest")
	if err != nil {
		rateStruct := newRepo.ReadLastest()
		j, err := json.Marshal(map[string]interface{}{
			"rate": rateStruct.ExchangeRate,
			"date": rateStruct.EffectiveDate,
		})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": "InternalServerError",
				},
			)
			return
		}

		rateJSON = string(j)
		redisRepo.Create("latest", string(j))

	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteString(rateJSON)
}

// GetHistoryRate 获取最近n天的历史汇率
func GetHistoryRate(c *gin.Context) {
	var form models.RateRequest
	err := c.BindJSON(&form)
	if err != nil {
		fmt.Println("form err", err)
		global.ErrorLogger.Println("[handlers]Form(GetHistoryRate) went wrong")
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
		fmt.Println("Unable to get the historical exchange rates that was queried")
		global.ErrorLogger.Println("Get the historical exchange rates went wrong")
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"historyRate": historyRate,
			},
		)
	}

}
