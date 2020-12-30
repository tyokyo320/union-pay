package global

import (
	"union-pay/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// 声明全局变量
var (
	CONFIG        *config.Config
	POSTGRESQL_DB *gorm.DB
	REDIS         *redis.Client
)
