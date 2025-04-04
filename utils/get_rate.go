package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"union-pay/global"
)

// 用于爬取
type Result struct {
	ExchangeRate []ExchangeRate `json:"exchangeRateJson"`
	CurDate      string         `json:"curDate"`
}

type ExchangeRate struct {
	TransactionCurrency string  `json:"transCur"`
	BaseCurrency        string  `json:"baseCur"`
	ExchangeRate        float64 `json:"rateData"`
}

// 银联官网爬取所选日期，货币种类的汇率
func GetRate(date string) (float64, error) {
	var d string
	if strings.Contains(date, "-") {
		d = strings.Replace(date, "-", "", -1)
	} else {
		d = date
	}
	urls := fmt.Sprintf("https://m.unionpayintl.com/jfimg/%s.json", d)
	fmt.Println(urls)

	client := &http.Client{}
	request, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		global.ErrorLogger.Println("[utils]Http get went wrong")
		fmt.Println("http get error", err)
		return 0.0, errors.New("http get error")
	}

	// 自定义Header
	// request.Header.Set("origin", "https://m.unionpayintl.com")
	// request.Header.Set("referer", "https://m.unionpayintl.com/cardholderServ/wap/rate?language=cn")
	// request.Header.Set("sec-fetch-dest", "empty")
	// request.Header.Set("sec-fetch-mode", "cors")
	// request.Header.Set("x-requested-with", "XMLHttpRequest")
	// request.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	request.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

	response, err := client.Do(request)
	if err != nil {
		global.ErrorLogger.Println("[utils]Http response went wrong")
		fmt.Println("response error", err)
		return 0.0, errors.New("response error")
	}

	// 函数结束后关闭相关链接
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		global.ErrorLogger.Println("[utils]Http read body went wrong")
		fmt.Println("read error", err)
		return 0.0, errors.New("read error")
	}
	// fmt.Println(string(body))

	res := Result{}
	json.Unmarshal(body, &res)
	// fmt.Println(res.CurDate)
	// fmt.Println(res.ExchangeRate)

	return res.ExchangeRate[387].ExchangeRate, nil
}
