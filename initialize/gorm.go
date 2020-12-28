package initialize

import (
	"fmt"
	"union-pay/config"
	"union-pay/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 依赖注入，连接数据库db依赖配置文件信息config
// (db -> config)
func NewGorm(c config.PostGreSQL) *gorm.DB {
	// 连接数据库
	// m := global.GVA_CONFIG
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		// m.PostGreSQL.Host,
		c.Host,
		c.User,
		c.Password,
		c.DB,
		c.Post,
		c.SSLmode,
		c.Timezone,
	)
	// fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection failure")
	}

	// Migrate the schema
	db.AutoMigrate(&models.TempRate{})
	db.AutoMigrate(&models.UpdateRate{})

	return db
}
