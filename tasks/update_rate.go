// 用于更新update DB，每天执行一次
package tasks

import (
	"fmt"
	"time"
	"union-pay/global"
	"union-pay/repository"
	"union-pay/utils"
)

// Job Specific Functions
type UpdateRate struct {
	// filtered
}

// 每天执行一次，如果当天没获取到，复制前一天有数据的汇率
func (e UpdateRate) Run() {
	// get lastest rate
	// date := "2020-12-28"
	currentTime := time.Now()
	date := currentTime.Format("2020-12-28")
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

	// update DB中先创建当天数据，并不断更新有变化的数据
	var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
	if isex, _ := newRepo.IsExist(date); isex {
		newRepo.Update(date, rate)
	} else {
		newRepo.Create(date, rate)
	}
	fmt.Println("------")

}
