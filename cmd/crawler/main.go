// 用于爬取历史数据
package main

import (
	"fmt"
	"time"

	"union-pay/config"
	"union-pay/global"
	"union-pay/initialize"
	"union-pay/repository"
	"union-pay/utils"
)

func init() {
	// Singleton pattern: 给全局变量赋值（给实例化的变量赋值）
	// 先读取配置文件
	global.CONFIG = config.NewConfig(".")
	// 然后才能初始化连接数据库
	global.POSTGRESQL_DB = initialize.NewGorm(global.CONFIG.PostGreSQL)
}

func main() {
	// set start of days
	start := time.Date(2020, 12, 1, 0, 0, 0, 0, time.UTC)
	end := time.Now()
	fmt.Println(start.Format("2006-01-02"), "-", end.Format("2006-01-02"))

	lastRate := 0.0

	// reverse
	for rd := rangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}

		// fmt.Println(date.Format("2006-01-02"))
		rate, err := utils.GetRate(date.Format("2006-01-02"), "CNY", "JPY")
		if err != nil {
			fmt.Println("[crawler]get history rate error!")
			global.ErrorLogger.Println("Get history rate went wrong")
			return
		}

		if rate == 0.0 {
			if lastRate == 0.0 {
				fmt.Println("first day no data")
				return
			}
			rate = lastRate
		} else {
			lastRate = rate
		}

		// add the history rates into update DB
		fmt.Println(date.Format("2006-01-02"), rate)
		// 如果当日还没更新，复制前一天不为0的数据
		var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
		newRepo.Create(date.Format("2006-01-02"), rate)

		time.Sleep(time.Second * 5)
	}
}

// rangeDate returns a date range function over start date to end date inclusive.
// After the end of the range, the range function returns a zero date,
// date.IsZero() is true.
func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}
