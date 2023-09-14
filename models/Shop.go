package models

import (
	"fmt"
	"time"

	"github.com/wagfog/hmdp_go/utils"
)

type Shop struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name"`
	TypeID     int64     `json:"typeId"`
	Images     string    `json:"images"`
	Area       string    `json:"area"`
	Address    string    `json:"address"`
	X          float64   `json:"x"`
	Y          float64   `json:"y"`
	AvgPrice   int64     `json:"avgPrice"`
	Sold       int       `json:"sold"`
	Comments   int       `json:"comments"`
	Score      int       `json:"score"`
	OpenHours  string    `json:"openHours"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Distance   float64   `gorm:"-"`
}

func QueryShopsByType(typeid int, current int) []Shop {
	var shops []Shop
	db.Table("tb_shop").Offset((current - 1) * utils.DEFAULT_PAGE_SIZE).Limit(utils.DEFAULT_PAGE_SIZE).Find(shops)
	return shops
}

func QueryShopsByIds(idStr string, idStrArr []string) []Shop {
	var shops []Shop
	db.Table("tb_shop").Where("id in (?)", idStrArr).Order("FIELD(id," + idStr + ")").Find(&shops)
	fmt.Println(shops)
	return shops
}

func QueryShopByid(id int64) *Shop {
	var shop Shop
	res := db.Table("tb_shop").Where("id = ?", id).First(&shop)
	if res.RecordNotFound() {
		return nil
	}
	return &shop
}
