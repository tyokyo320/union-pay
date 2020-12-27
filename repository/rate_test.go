package repository

import (
	"fmt"
	"testing"
	"union-pay/config"
	"union-pay/global"
	"union-pay/initialize"
	"union-pay/models"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	// 为了测试，必须先读取配置文件
	global.CONFIG = config.NewConfig("../.")
	// 然后连接数据库
	db := initialize.NewGorm(global.CONFIG.PostGreSQL)

	instances := []models.Rate{
		{
			// ExchangeRateID:      2413414,
			// CurDate:             time.Unix(1608739200, 0),
			BaseCurrency:        "CNY",
			TransactionCurrency: "JPY",
			ExchangeRate:        0.063361,
			// CreateDate:          time.Unix(1608739200, 0),
			// UpdateDate:          time.Unix(1608739200, 0),
			EffectiveDate: "2020-12-24",
			Time:          "12:45",
		},
		{
			// ExchangeRateID:      2413414,
			// CurDate:             time.Unix(1608739200, 0),
			BaseCurrency:        "CNY",
			TransactionCurrency: "JPY",
			ExchangeRate:        0.06666,
			// CreateDate:          time.Unix(1608739200, 0),
			// UpdateDate:          time.Unix(1608739200, 0),
			EffectiveDate: "2020-12-24",
			Time:          "14:37",
		},
		// {
		// 	ExchangeRateID:      2411188,
		// 	CurDate:             time.Unix(1608652800, 0),
		// 	BaseCurrency:        "CNY",
		// 	TransactionCurrency: "JPY",
		// 	ExchangeRate:        0.063515,
		// 	CreateDate:          time.Unix(1608652800, 0),
		// 	UpdateDate:          time.Unix(1608652800, 0),
		// 	EffectiveDate:       time.Unix(1608711868, 0),
		// },
		// {
		// 	ExchangeRateID:      2408962,
		// 	CurDate:             time.Unix(1608566400, 0),
		// 	BaseCurrency:        "CNY",
		// 	TransactionCurrency: "JPY",
		// 	ExchangeRate:        0.06363,
		// 	CreateDate:          time.Unix(1608566400, 0),
		// 	UpdateDate:          time.Unix(1608566400, 0),
		// 	EffectiveDate:       time.Unix(1608625456, 0),
		// },
		// {
		// 	ExchangeRateID:      2404510,
		// 	CurDate:             time.Unix(1608220800, 0),
		// 	BaseCurrency:        "CNY",
		// 	TransactionCurrency: "JPY",
		// 	ExchangeRate:        0.063523,
		// 	CreateDate:          time.Unix(1608220800, 0),
		// 	UpdateDate:          time.Unix(1608220800, 0),
		// 	EffectiveDate:       time.Unix(1608279476, 0),
		// },
	}

	// Create
	result := db.Create(&instances[1])

	require.NoError(t, result.Error)
}

// func TestDelete(t *testing.T) {
// }

func TestGroup(t *testing.T) {
	// 为了测试，必须先读取配置文件
	global.CONFIG = config.NewConfig("../.")
	// 然后连接数据库
	db := initialize.NewGorm(global.CONFIG.PostGreSQL)

	page := 1
	pageSize := 10

	subQuery := db.
		Table("rates").
		Select("rates.effective_date, MAX(rates.time) as max_time").
		Group("effective_date")

	rows, err := db.Table("(?) as u", db.Model(&models.Rate{})).
		Select("u.effective_date, u.time, u.exchange_rate").
		Joins(
			"JOIN (?) as v on u.effective_date = v.effective_date AND u.time = v.max_time",
			subQuery,
		).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Rows()

	if err != nil {
		t.Fail()
	}

	// t.Log("rows")

	var date string
	var time string
	var exchangeRate float64

	for rows.Next() {
		rows.Scan(&date, &time, &exchangeRate)
		fmt.Println(date, time, exchangeRate)
	}
}
