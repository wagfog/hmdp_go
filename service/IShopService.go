package service

import (
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
)

type IShopService interface {
	QueryById(id int64) result.Result

	Update(shop models.Shop) result.Result

	QueryShopByType(typeId int, current int, x float64, y float64) result.Result
}
