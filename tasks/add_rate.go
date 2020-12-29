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
	// date := "2020-12-28"
	// time := "19:05:12"
	currentTime := time.Now()
	date := currentTime.Format("2020-12-28")
	time := currentTime.Format("19:05:12")
	// fmt.Println("Current Time in String: ", currentTime.String())
	rate, err := utils.GetRate(date, "CNY", "JPY")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(rate)
	if rate == 0 {
		fmt.Println("当日汇率查询显示待更新")
		return
	}

	// 在temp rate DB中一直添加最新爬取的数据
	var repo *repository.RateRepository = repository.NewRateRepository(global.POSTGRESQL_DB)
	err = repo.Create(date, time, rate)
	if err != nil {
		fmt.Println("temp rate添加数据失败")
		return
	}

	// 检查update数据库中是否有数据，有则更新，没有插入
	var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
	if isex, _ := newRepo.IsExist(date); isex {
		newRepo.Update(date, rate)
	} else {
		newRepo.Create(date, rate)
	}

	// Sends some email
	// fmt.Printf("Every 10 sec send reminder emails \n")
}
