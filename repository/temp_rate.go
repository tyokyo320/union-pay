package repository

import (
	"errors"
	"fmt"
	"union-pay/models"

	"gorm.io/gorm"
)

type RateHistory struct {
	Date         string  `json:"date"`
	Time         string  `json:"time"`
	ExchangeRate float64 `json:"exchangeRate"`
}

// 定义接口
type IRepository interface {
	Create(date, time string, rate float64) error
	Read(date string) *models.TempRate
	ReadLastest() *models.TempRate
	ReadList(page, pageSize int) ([]RateHistory, error)
	// Update(date string, newdata float64)
	// Delete(id uint)
}

// 定义结构体
type RateRepository struct {
	db *gorm.DB
}

// 依赖注入，repository依赖连接数据库db
// (repository -> db)
func NewRateRepository(db *gorm.DB) *RateRepository {
	// 返回一个repository的实例
	return &RateRepository{db}
}

// 实现接口方法
// 向TempRate数据库添加数据
func (r *RateRepository) Create(date, time string, rate float64) error {
	add := []models.TempRate{
		{
			BaseCurrency:        "CNY",
			TransactionCurrency: "JPY",
			ExchangeRate:        rate,
			EffectiveDate:       date,
			Time:                time,
		},
	}
	result := r.db.Create(&add)

	if result.Error != nil {
		return errors.New("Create Temp Rate error")
	}

	return nil
}

// 从数据库读取所选择的具体某一天的汇率
func (r *RateRepository) Read(date string) *models.TempRate {
	rate := models.TempRate{}
	result := r.db.Where("effective_date = ?", date).Order("time DESC").First(&rate)
	if result.Error != nil {
		return nil
	}

	return &rate
}

// func (r *RateRepository) Update(date string, newdata float64) {
// 	updateRate := models.UpdateRate{}

// 	r.db.Model(&updateRate).Where("effective_date = ?", date).Update("exchange_rate", newdata)
// }

// func (r *RateRepository) Delete(id uint) {
// 	rate := models.TempRate{}
// 	r.db.Delete(&rate, id)
// }

// 从数据库读取最近一天的汇率
func (r *RateRepository) ReadLastest() *models.TempRate {
	lastestRate := models.TempRate{}
	result := r.db.Order("effective_date DESC").Order("time DESC").First(&lastestRate)
	if result.Error != nil {
		return nil
	}

	return &lastestRate
}

// 从数据库获取近n天的历史汇率
func (r *RateRepository) ReadList(page, pageSize int) ([]RateHistory, error) {
	if page < 1 {
		return nil, errors.New("page < 1")
	}

	subQuery := r.db.
		Table("temp_rates").
		Select("rates.effective_date, MAX(rates.time) as max_time").
		Group("effective_date")

	rows, err := r.db.Table("(?) as u", r.db.Model(&models.TempRate{})).
		Select("u.effective_date, u.time, u.exchange_rate").
		Joins(
			"JOIN (?) as v on u.effective_date = v.effective_date AND u.time = v.max_time",
			subQuery,
		).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Rows()

	if err != nil {
		return nil, errors.New("query error")
	}

	history := []RateHistory{}

	for rows.Next() {
		h := RateHistory{}
		rows.Scan(&h.Date, &h.Time, &h.ExchangeRate)
		fmt.Println(h.Date, h.Time, h.ExchangeRate)
		history = append(history, h)
	}

	return history, nil

}
