package repository

import (
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
	result := db.Create(&instances[0])

	require.NoError(t, result.Error)
}

// func TestDelete(t *testing.T) {
// }
