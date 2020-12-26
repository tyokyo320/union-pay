package repository

import (
	"union-pay/models"

	"gorm.io/gorm"
)

// 定义接口
type IRepository interface {
	Create(rate *models.Rate) (uint, error)
	Read(date string) *models.Rate
	Update(id uint, new *models.Rate)
	Delete(id uint)
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
func (r *RateRepository) Create(rate *models.Rate) (uint, error) {
	result := r.db.Create(&rate)
	return rate.ID, result.Error
}

func (r *RateRepository) Read(date string) *models.Rate {
	rate := models.Rate{}
	r.db.Where("effective_date = ?", date).Order("time DESC").First(&rate)
	return &rate
}

func (r *RateRepository) Delete(id uint) {
	rate := models.Rate{}
	r.db.Delete(&rate, id)
}
