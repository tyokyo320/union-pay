package config

import (
	"os"

	"github.com/spf13/viper"
)

// NewConfig　用于读取配置文件来连接数据库
func NewConfig(path string) *Config {
	// 配置文件所在路径
	viper.AddConfigPath(path)
	// 配置文件名
	if os.Getenv("GIN_MODE") == "release" {
		viper.SetConfigName("config")
	} else {
		viper.SetConfigName("devconfig")
	}
	// 配置文件名后缀
	viper.SetConfigType("yml")

	// 读取配置文件的内容
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 实例化
	config := Config{}

	// 赋值给config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
