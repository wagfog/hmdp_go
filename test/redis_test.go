package test

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis"
	"github.com/wagfog/hmdp_go/config/setting"
)

func TestRedis(t *testing.T) {
	setting.Init()
	Client := redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
	})

	val, err := Client.Get("tmpkey").Result()
	if err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Println("the val is :", val)

}
