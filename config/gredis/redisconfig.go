package gredis

import (
	"github.com/go-redis/redis/v8"
	"github.com/wagfog/hmdp_go/config/setting"
)

var Client *redis.Client

func Setup() {
	Client = redis.NewClient(&redis.Options{
		Addr:        setting.RedisSetting.Host,
		Password:    setting.RedisSetting.Password,
		PoolSize:    setting.RedisSetting.MaxIdle,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
	})
}
