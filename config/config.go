package config

// 用于保存各种数据库
type Config struct {
	PostGreSQL PostGreSQL `mapstructure:"postgresql"`
	Redis      Redis      `mapstructure:"redis"`
}

// Postgresql
type PostGreSQL struct {
	Host     string `mapstructure:"host"`
	Post     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
	SSLmode  string `mapstructure:"sslmode"`
	Timezone string `mapstructure:"timezone"`
}

// Redis
type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
