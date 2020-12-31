// 用于更新temp DB，每天执行多次，不断储存获取的当天数据
package tasks

import (
	"fmt"
	"time"
	"union-pay/global"
	"union-pay/repository"
	"union-pay/utils"
)

// Job Specific Functions
type AddRate struct {
	// filtered
}

// ReminderEmails.Run() will get triggered automatically.
func (e AddRate) Run() {
	// get lastest rate
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	time := currentTime.Format("15:04:05")
	rate, err := utils.GetRate(date, "CNY", "JPY")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(rate)
	if rate == 0 {
		fmt.Println("当日汇率查询显示待更新!")
		return
	}

	// 在temp_rate DB中一直添加最新爬取的数据
	var repo *repository.RateRepository = repository.NewRateRepository(global.POSTGRESQL_DB)
	err = repo.Create(date, time, rate)
	if err != nil {
		fmt.Println("temp rate添加数据失败")
		return
	}

	// 检查update数据库中是否有数据，有更新则更新DB，没有插入
	var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
	if change := newRepo.Read(date); change != nil {
		if change.ExchangeRate != rate {
			newRepo.Update(date, rate)
		}
	} else {
		newRepo.Create(date, rate)
	}
	// 并添加至缓存中
	var redisRepo *repository.RateCacheRepository = repository.NewRateCacheRepository(global.REDIS)
	redisRepo.Create("latest", rate)

	// Sends some email
	// fmt.Printf("Every 10 sec send reminder emails \n")
}
