package repository

import (
	"errors"
	"union-pay/models"

	"gorm.io/gorm"
)

// 定义接口
type IUpdateRepository interface {
	CreateUpdateRate(date string, rate float64) error
	Read(date string) *models.TempRate
	Update(date string, newdata float64)
}

// 定义结构体
type UpdateRateRepository struct {
	db *gorm.DB
}

// 依赖注入，repository依赖连接数据库db
// (repository -> db)
func NewUpdateRateRepository(db *gorm.DB) *UpdateRateRepository {
	// 返回一个repository的实例
	return &UpdateRateRepository{db}
}

// 向UpdateRate数据库添加数据
func (r *UpdateRateRepository) CreateUpdateRate(date string, rate float64) error {
	add := models.UpdateRate{
		BaseCurrency:        "CNY",
		TransactionCurrency: "JPY",
		ExchangeRate:        rate,
		EffectiveDate:       date,
	}

	result := r.db.Create(&add)

	if result.Error != nil {
		return errors.New("Create Update Rate error")
	}

	return nil
}

func (r *UpdateRateRepository) Update(date string, newdata float64) {
	updateRate := models.UpdateRate{}

	r.db.Model(&updateRate).Where("effective_date = ?", date).Update("exchange_rate", newdata)
}
