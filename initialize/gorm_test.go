package initialize

import (
	"testing"
	"union-pay/config"
	"union-pay/global"
)

func TestDatabaseConnection(t *testing.T) {
	// 为了测试，必须先读取配置文件
	global.CONFIG = config.NewConfig("../.")
	// 然后连接数据库
	NewGorm(global.CONFIG.PostGreSQL)
}
