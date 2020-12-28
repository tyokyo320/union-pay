package models

import (
	"gorm.io/gorm"
)

type UpdateRate struct {
	gorm.Model
	BaseCurrency        string  `json:"baseCurrency"`
	TransactionCurrency string  `json:"transactionCurrency"`
	ExchangeRate        float64 `json:"exchangeRate"`
	EffectiveDate       string  `json:"effectiveDate"`
}
