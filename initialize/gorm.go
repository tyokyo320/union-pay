package initialize

import (
	"fmt"
	"union-pay/global"
	"union-pay/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Gorm init database connection
func Gorm() *gorm.DB {
	// 连接数据库
	// m := global.GVA_CONFIG
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		// m.PostGreSQL.Host,
		global.GVA_CONFIG.PostGreSQL.Host,
		global.GVA_CONFIG.PostGreSQL.User,
		global.GVA_CONFIG.PostGreSQL.Password,
		global.GVA_CONFIG.PostGreSQL.DB,
		global.GVA_CONFIG.PostGreSQL.Post,
		global.GVA_CONFIG.PostGreSQL.SSLmode,
		global.GVA_CONFIG.PostGreSQL.Timezone,
	)
	// fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection failure")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Rate{})

	return db
}
