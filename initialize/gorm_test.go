package initialize

import (
	"testing"
	"union-pay/config"
	"union-pay/global"
	"union-pay/models"

	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	// 为了测试，必须先读取配置文件
	global.CONFIG = config.NewConfig("../.")
	// 然后连接数据库
	NewGorm(global.CONFIG.PostGreSQL)
}

func TestRead(t *testing.T) {
	global.CONFIG = config.NewConfig("../.")
	db := NewGorm(global.CONFIG.PostGreSQL)

	rate := models.TempRate{}
	var exchangeRateID uint = 2413414

	// Read
	db.First(&rate, "exchange_rate_id = ?", exchangeRateID)
	require.Equal(t, rate.ExchangeRateID, exchangeRateID)
}
