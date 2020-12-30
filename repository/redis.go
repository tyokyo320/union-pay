package repository

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// 定义redis接口
type IRateCacheRepository interface {
	Create(date string, rate float64) error
	Read(date string) (float64, error)
}

// 定义结构体
type RateCacheRepository struct {
	rdb *redis.Client
}

// 创建一个结构体实例
func NewRateCacheRepository(rdb *redis.Client) *RateCacheRepository {
	return &RateCacheRepository{rdb}
}

// 在缓存中创建
func (r *RateCacheRepository) Create(date string, rate float64) error {
	err := r.rdb.Set(context.TODO(), date, rate, 0).Err()
	return err
}

// 从缓存中读取某一天汇率
func (r *RateCacheRepository) Read(date string) (float64, error) {
	result, err := r.rdb.Get(context.TODO(), date).Result()
	if err != nil {
		return 0.0, err
	}
	// 转float
	return strconv.ParseFloat(result, 64)
}
