package main

import (
	"union-pay/config"
	"union-pay/global"
	"union-pay/initialize"
)

func main() {
	// 给全局变量赋值
	global.GVA_CONFIG = config.LodeConfig()
	global.GVA_DB = initialize.Gorm()

}
