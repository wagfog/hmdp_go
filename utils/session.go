package utils

import (
	"github.com/gin-contrib/sessions/redis"
	"github.com/wagfog/hmdp_go/config/setting"
)

var (
	Redistore redis.Store
	err       error
)

func InitRedistore() {
	Redistore, err = redis.NewStore(setting.RedisSetting.MaxIdle, "tcp", setting.RedisSetting.Host, setting.RedisSetting.Password, []byte("swag-secret"))
	if err != nil {
		panic(err)
	}
}
