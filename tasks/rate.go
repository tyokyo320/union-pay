package tasks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"union-pay/global"
	"union-pay/repository"

	"github.com/bamzi/jobrunner"
)

// 用于爬取
type result struct {
	ExchangeRateId            int
	CurDate                   int
	BaseCurrency              string
	TransactionCurrency       string
	ExchangeRate              float64
	CreateDate                int
	CreateUser                int
	UpdateDate                int
	UpdateUser                int
	EffectiveDate             int
	TransactionCurrencyOption interface{}
}

func RunTasks() {
	// optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Start()
	jobrunner.Schedule("@every 10s", AddRate{})
	jobrunner.Schedule("@every 30s", UpdateRate{})
}

// Job Specific Functions
type AddRate struct {
	// filtered
}

// ReminderEmails.Run() will get triggered automatically.
func (e AddRate) Run() {
	// get lastest rate
	currentTime := time.Now()
	date := currentTime.Format("2020-12-28")
	// date := "2020-11-19"
	time := currentTime.Format("19:05:12")
	// fmt.Println("Current Time in String: ", currentTime.String())
	rate, err := GetRate(date, "CNY", "JPY")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(rate)
		if rate == 0 {
			fmt.Println("当日汇率查询显示待更新")
			return
		}
		// 在temp rate DB中一直添加最新爬取的数据
		var repo *repository.RateRepository = repository.NewRateRepository(global.POSTGRESQL_DB)
		err := repo.CreateTempRate(date, time, rate)
		if err != nil {
			fmt.Println("temp rate添加数据失败")
			return
		}
		// 检查update数据库中是否有数据，有则更新，没有插入

	}

	// Sends some email
	// fmt.Printf("Every 10 sec send reminder emails \n")
}

// Job Specific Functions
type UpdateRate struct {
	// filtered
}

// 每天执行一次，如果当天没获取到，复制前一天有数据的汇率
func (e UpdateRate) Run() {
	// get lastest rate
	currentTime := time.Now()
	date := currentTime.Format("2020-12-28")
	// date := "2020-11-19"
	rate, err := GetRate(date, "CNY", "JPY")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(rate)
		if rate == 0 {
			fmt.Println("当日汇率查询显示待更新")
			return
		}

		// update DB中先创建当天数据，并不断更新有变化的数据
		var newRepo *repository.UpdateRateRepository = repository.NewUpdateRateRepository(global.POSTGRESQL_DB)
		err = newRepo.CreateUpdateRate(date, rate)
		if err != nil {
			fmt.Println("update rate添加数据失败")
			return
		}
		newRepo.Update(date, rate)
		fmt.Println("------")
	}

}

// 银联官网爬取所选日期，货币种类的汇率
func GetRate(date, baseCurrency, transactionCurrency string) (float64, error) {
	client := &http.Client{}

	// Form data
	data := url.Values{}
	data.Add("curDate", date)
	data.Add("baseCurrency", baseCurrency)
	data.Add("transactionCurrency", transactionCurrency)

	request, err := http.NewRequest("POST", "https://m.unionpayintl.com/cardholderServ/wap/rate/search", bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("http post error", err)
		return 0.0, errors.New("http post error")
	}

	// 自定义Header
	request.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("response error", err)
		return 0.0, errors.New("response error")
	}

	// 函数结束后关闭相关链接
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read error", err)
		return 0.0, errors.New("read error")
	}
	// fmt.Println(string(body))

	// convert JSON to struct
	res := result{}
	json.Unmarshal(body, &res)
	// fmt.Println(res)
	// fmt.Println(res.ExchangeRate)

	return res.ExchangeRate, nil
}
