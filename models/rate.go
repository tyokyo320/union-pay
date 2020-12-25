package models

import "time"

type Rate struct {
	ExchangeRateID      uint      `json:"exchangeRateId"`
	CurDate             time.Time `json:"curDate"`
	BaseCurrency        string    `json:"baseCurrency"`
	TransactionCurrency string    `json:"transactionCurrency"`
	ExchangeRate        float64   `json:"exchangeRate"`
	CreateDate          time.Time `json:"createDate"`
	UpdateDate          time.Time `json:"updateDate"`
	EffectiveDate       time.Time `json:"effectiveDate"`
}
