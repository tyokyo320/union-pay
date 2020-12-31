package models

// RateRequest form request
type RateRequest struct {
	Date     string `json:"date"`
	Page     int    `json:"page" binding:"gte=1"`
	PageSize int    `json:"pageSize" binding:"gte=1,lte=50"`
}
