package initialize

import (
	"union-pay/config"

	"github.com/go-redis/redis/v8"
)

// NewRedis connection
func NewRedis(c config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password, // no password set
		DB:       c.DB,       // use default DB
	})

	return client
}
