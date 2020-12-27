package repository

import (
	"union-pay/models"

	"gorm.io/gorm"
)

// 定义接口
type IRepository interface {
	Create(rate *models.Rate) (uint, error)
	Read(date string) *models.Rate
	ReadLastest() *models.Rate
	ReadHistory(n int) []*models.Rate
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

// 读取所选择的具体某一天的汇率
func (r *RateRepository) Read(date string) *models.Rate {
	rate := models.Rate{}
	result := r.db.Where("effective_date = ?", date).Order("time DESC").First(&rate)
	if result.Error != nil {
		return nil
	}

	return &rate
}

func (r *RateRepository) Delete(id uint) {
	rate := models.Rate{}
	r.db.Delete(&rate, id)
}

// 读取最近一天的汇率
func (r *RateRepository) ReadLastest() *models.Rate {
	lastestRate := models.Rate{}
	result := r.db.Order("effective_date DESC").Order("time DESC").First(&lastestRate)
	if result.Error != nil {
		return nil
	}

	return &lastestRate
}

func (r *RateRepository) ReadHistory(n int) []*models.Rate {
	if n < 1 {
		panic("n should be a positive integer greater than 0")
	}

	historyRate := []models.Rate{}

}
