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

func TestRedis(t *testing.T) {
	setting.Init()
	Client := redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
	})

	val, err := Client.Get(utils.CACHE_SHOP_KEY + "1").Result()
	if err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Println(val)
	var shop models.Shop
	err = json.Unmarshal([]byte(val), &shop)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("ID: %d\n", shop.ID)
	fmt.Printf("Name: %s\n", shop.Name)
	fmt.Printf("TypeID: %d\n", shop.TypeID)
	fmt.Printf("Images: %s\n", shop.Images)
	fmt.Printf("Area: %s\n", shop.Area)
	fmt.Printf("Address: %s\n", shop.Address)
	fmt.Printf("X: %f\n", shop.X)
	fmt.Printf("Y: %f\n", shop.Y)
	fmt.Printf("AvgPrice: %d\n", shop.AvgPrice)
	fmt.Printf("Sold: %d\n", shop.Sold)
	fmt.Printf("Comments: %d\n", shop.Comments)
	fmt.Printf("Score: %d\n", shop.Score)
	fmt.Printf("OpenHours: %s\n", shop.OpenHours)
	fmt.Printf("CreateTime: %s\n", shop.CreateTime)
	fmt.Printf("UpdateTime: %s\n", shop.UpdateTime)
	fmt.Printf("Distance: %f\n", shop.Distance)

}
