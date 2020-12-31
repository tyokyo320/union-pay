package models

import (
	"gorm.io/gorm"
)

// TempRate defined temp_rate DB table
type TempRate struct {
	gorm.Model
	// ExchangeRateID      uint      `json:"exchangeRateId"`
	// CurDate             time.Time `json:"curDate"`
	BaseCurrency        string  `json:"baseCurrency"`
	TransactionCurrency string  `json:"transactionCurrency"`
	ExchangeRate        float64 `json:"exchangeRate"`
	// CreateDate          time.Time `json:"createDate"`
	// UpdateDate          time.Time `json:"updateDate"`
	EffectiveDate string `json:"effectiveDate"`
	Time          string `json:"time"`
}

// TempRateHistory defined temp_rate DB history rate
type TempRateHistory struct {
	Date         string  `json:"date"`
	Time         string  `json:"time"`
	ExchangeRate float64 `json:"exchangeRate"`
}
