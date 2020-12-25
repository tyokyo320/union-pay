package global

import (
	"union-pay/config"

	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG *config.Config
)
