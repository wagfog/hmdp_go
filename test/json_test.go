package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-redis/redis"
	"github.com/wagfog/hmdp_go/config/setting"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/utils"
)

func TestJson(t *testing.T) {
	setting.Init()
	models.Init()
	Client := redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
	})
	blog := models.QueryShopByid(2)
	// fmt.Println(blog)
	json_, er := json.Marshal(blog)
	if er != nil {
		fmt.Println(er)
		return
	}
	fmt.Println(string(json_))
	Client.Set(utils.CACHE_SHOP_KEY+"2", string(json_), 0)
}
