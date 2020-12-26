package global

import (
	"union-pay/config"

	"gorm.io/gorm"
)

// 声明全局变量
var (
	CONFIG        *config.Config
	POSTGRESQL_DB *gorm.DB
)
