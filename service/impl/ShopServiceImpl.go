package impl

import (
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
)

type ShopServiceImpl struct {
}

func (shopServiceImpl *ShopServiceImpl) QueryById(id int64) result.Result {
	return *result.Fail("not finish")
}

func (shopServiceImpl *ShopServiceImpl) Update(shop models.Shop) result.Result {
	return *result.Fail("not finish")
}

func (shopServiceImpl *ShopServiceImpl) QueryShopByType(typeId int, current int, x float64, y float64) result.Result {
	return *result.Fail("not finish")
}
