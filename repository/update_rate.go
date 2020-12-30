package repository

import (
	"errors"
	"union-pay/models"

	"gorm.io/gorm"
)

// 定义接口
type IUpdateRepository interface {
	Create(date string, rate float64) error
	Read(date string) *models.TempRate
	ReadLastest() *models.TempRate
	ReadList(page, pageSize int) ([]models.UpdateRate, error)
	Update(date string, rate float64)
	// IsExist(date string) (bool, error)
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
func (r *UpdateRateRepository) Create(date string, rate float64) error {
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

// 从Update DB 数据库读取所选择的具体某一天的汇率
func (r *UpdateRateRepository) Read(date string) *models.UpdateRate {
	rate := models.UpdateRate{}
	result := r.db.Where("effective_date = ?", date).Order("effective_date DESC").First(&rate)
	if result.Error != nil {
		return nil
	}

	return &rate
}

// 从Update DB 数据库读取最近一天的汇率
func (r *UpdateRateRepository) ReadLastest() *models.UpdateRate {
	lastestRate := models.UpdateRate{}
	result := r.db.Order("effective_date DESC").First(&lastestRate)
	if result.Error != nil {
		return nil
	}

	return &lastestRate
}

// 从Update DB数据库获取近n天的历史汇率
func (r *UpdateRateRepository) ReadList(page, pageSize int) ([]models.UpdateRate, error) {
	var history []models.UpdateRate

	result := r.db.Model(&models.UpdateRate{}).
		Order("effective_date DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&history)

	if result.Error != nil {
		return nil, errors.New("query Update DB history rate error")
	}

	return history, nil
}

// 判断当天数据是否存在
// func (r *UpdateRateRepository) IsExist(date string) (bool, error) {
// 	updateRate := models.UpdateRate{}
// 	if err := r.db.Where("effective_date = ?", date).Order("effective_date DESC").First(&updateRate).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			fmt.Println("Error! Record Not Found")
// 			return false, nil
// 		}
// 		return false, err
// 	}

// 	return true, nil
// }

// 用于更新update DB数据库当日汇率
func (r *UpdateRateRepository) Update(date string, rate float64) {
	updateRate := models.UpdateRate{}
	r.db.Model(updateRate).Where("effective_date = ?", date).Order("effective_date DESC").Update("exchange_rate", rate)
}
