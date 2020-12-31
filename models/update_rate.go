package models

import (
	"gorm.io/gorm"
)

// UpdateRate defined update_rate DB table
type UpdateRate struct {
	gorm.Model
	BaseCurrency        string  `json:"baseCurrency" gorm:"uniqueIndex:idx_name"`
	TransactionCurrency string  `json:"transactionCurrency" gorm:"uniqueIndex:idx_name"`
	ExchangeRate        float64 `json:"exchangeRate"`
	EffectiveDate       string  `json:"effectiveDate" gorm:"uniqueIndex:idx_name"`
}
