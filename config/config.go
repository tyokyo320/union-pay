package config

import "github.com/spf13/viper"

type Config struct {
	PostGreSQL PostGreSQL `mapstructure:"postgresql"`
}

type PostGreSQL struct {
	Host     string `mapstructure:"host"`
	Post     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
	SSLmode  string `mapstructure:"sslmode"`
	Timezone string `mapstructure:"timezone"`
}

// 用于读取配置文件来连接数据库
func NewConfig(path string) *Config {
	// 配置文件所在路径
	viper.AddConfigPath(path)
	// 配置文件名
	viper.SetConfigName("config")
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
