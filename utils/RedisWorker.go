package utils

import (
	"context"
	"time"

	"github.com/wagfog/hmdp_go/config/gredis"
)

const (
	//开始时间戳
	BEGIN_TIMESTAMP int64 = 1640995200
	//序列号，位数
	COUNT_BITS int64 = 32
)

// type RedisWorker struct {

// }

func NextId(keyPrefix string) int64 {
	//1.生成时间戳
	now := time.Now().UTC()
	nowSecond := now.Unix()
	timestamp := nowSecond - BEGIN_TIMESTAMP

	//2.生成序列号
	//获取当前日期，精确到天
	date := now.Format("2006:01:02")
	//自增长
	count, _ := gredis.Client.Incr(context.Background(), "icr:"+keyPrefix+date).Result()
	//3.拼接并返回
	return (timestamp << COUNT_BITS) | count //实际上就是前32位是时间，后面是计数
}
