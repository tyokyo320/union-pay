package global

import (
	"log"
	"union-pay/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// 声明全局变量
var (
	CONFIG        *config.Config
	POSTGRESQL_DB *gorm.DB
	REDIS         *redis.Client
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	PanicLogger   *log.Logger
)
