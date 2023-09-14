package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/wagfog/hmdp_go/models"
)

func TestJWT(t *testing.T) {
	jsonData := `{"area":"大关","openHours":"10:00-22:00","sold":4215,"images":"https://web-tlias-swag.oss-cn-fuzhou.aliyuncs.com/9217d8b0-5446-4e52-9d9a-fa5ff3c77fe7.jpg","address":"金华路锦昌文华苑29号","comments":3035,"avgPrice":80,"updateTime":1690683185000,"score":37,"createTime":1640167839000,"name":"大歪哥餐厅","x":120.149192,"y":30.316078,"typeId":1,"id":1}`

	var shop models.Shop
	err := json.Unmarshal([]byte(jsonData), &shop)
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
