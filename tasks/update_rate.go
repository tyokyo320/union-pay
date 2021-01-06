// 用于更新update DB，每天执行一次
package tasks

import (
	"encoding/json"
	"time"
	"union-pay/global"
	"union-pay/repository"
)

// Job Specific Functions
type UpdateRate struct {
	// filtered
}

// 每天执行一次，如果当天没获取到，复制前一天有数据的汇率，
// 防止当天没有数据导致当天数据为空
func (e UpdateRate) Run() {
	global.InfoLogger.Println("[tasks Update]Job runner started...")
	// get lastest rate
	// date := "2020-12-28"
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")

	// 检查update数据库中是否有当天数据
	var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
	if rate := newRepo.Read(date); rate != nil {
		return
	}

	yesterdayTime := currentTime.AddDate(0, 0, -1)
	yesterday := yesterdayTime.Format("2006-01-02")

	rate := newRepo.Read(yesterday)
	if rate != nil {
		// 若当天无数据，则复制昨天数据为当天汇率
		newRepo.Create(date, rate.ExchangeRate)
	} else {
		// 若昨天数据为空，则panic
		global.PanicLogger.Println("[tasks update]Copy yesterday rate to update_rate DB went wrong")
		panic("No rate yesterday")
	}

	// 将数据编码成json字符串
	j, err := json.Marshal(map[string]interface{}{
		"rate": rate.ExchangeRate,
		"date": rate.EffectiveDate,
	})
	if err != nil {
		global.ErrorLogger.Println("[tasks update]Json marshal went wrong")
		return
	}

	// 并添加至缓存中
	var redisRepo *repository.RateCacheRepository = repository.NewRateCacheRepository(global.REDIS)
	redisRepo.Create("latest", string(j))
}
