package initialize

import (
	"testing"
	"time"
	"union-pay/models"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	db := Gorm()

	instances := []models.Rate{
		{
			ExchangeRateID:      2413414,
			CurDate:             time.Unix(1608739200, 0),
			BaseCurrency:        "CNY",
			TransactionCurrency: "JPY",
			ExchangeRate:        0.063361,
			CreateDate:          time.Unix(1608739200, 0),
			UpdateDate:          time.Unix(1608739200, 0),
			EffectiveDate:       time.Unix(1608798082, 0),
		},
		{
			ExchangeRateID:      2411188,
			CurDate:             time.Unix(1608652800, 0),
			BaseCurrency:        "CNY",
			TransactionCurrency: "JPY",
			ExchangeRate:        0.063515,
			CreateDate:          time.Unix(1608652800, 0),
			UpdateDate:          time.Unix(1608652800, 0),
			EffectiveDate:       time.Unix(1608711868, 0),
		},
	}

	// Create
	result := db.Create(&instances[0])

	require.NoError(t, result.Error)
}

func TestRead(t *testing.T) {
	db := Gorm()

	rate := models.Rate{}
	var exchangeRateID uint = 2413414

	// Read
	db.First(&rate, "exchange_rate_id = ?", exchangeRateID)
	require.Equal(t, rate.ExchangeRateID, exchangeRateID)
}
