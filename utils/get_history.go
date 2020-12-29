package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Historyresult struct {
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

func main() {
	start := time.Now()
	end := start.AddDate(0, 0, 10)
	fmt.Println(start.Format("2006-01-02"), "-", end.Format("2006-01-02"))

	for rd := rangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		fmt.Println(date.Format("2006-01-02"))
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

func getRate(date string) (float64, error) {
	client := &http.Client{}

	// Form data
	data := url.Values{}
	data.Add("curDate", date)
	data.Add("baseCurrency", "CNY")
	data.Add("transactionCurrency", "JPY")

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
	res := Historyresult{}
	json.Unmarshal(body, &res)
	// fmt.Println(res)
	// fmt.Println(res.ExchangeRate)

	return res.ExchangeRate, nil
}
