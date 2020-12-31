package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 用于爬取
type result struct {
	ExchangeRateID            int     `json:"exchangeRateId"`
	CurDate                   int     `json:"curDate"`
	BaseCurrency              string  `json:"baseCurrency"`
	TransactionCurrency       string  `json:"transactionCurrency"`
	ExchangeRate              float64 `json:"exchangeRate"`
	CreateDate                int     `json:"createDate"`
	CreateUser                int     `json:"createUser"`
	UpdateDate                int     `json:"updateDate"`
	UpdateUser                int     `json:"updateUser"`
	EffectiveDate             int     `json:"effectiveDate"`
	TransactionCurrencyOption interface{}
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
	request.Header.Set("origin", "https://m.unionpayintl.com")
	request.Header.Set("referer", "https://m.unionpayintl.com/cardholderServ/wap/rate?language=cn")
	request.Header.Set("sec-fetch-dest", "empty")
	request.Header.Set("sec-fetch-mode", "cors")
	request.Header.Set("x-requested-with", "XMLHttpRequest")
	request.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
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
