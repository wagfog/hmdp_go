package impl

import (
	"context"
	"encoding/json"
	"time"

	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/utils"
)

type ShopTypeService struct {
}

func NewShopTypeService() *ShopTypeService {
	return &ShopTypeService{}
}

func (shopTypeService *ShopTypeService) QueryAllList() result.Result {
	ans := make([]models.ShopType, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	shop_type_json := gredis.Client.SMembers(ctx, utils.SHOP_TYPE_KEY).Val()

	if shop_type_json != nil {
		var shop_type models.ShopType
		for _, j := range shop_type_json {
			json.Unmarshal([]byte(j), &shop_type)
			ans = append(ans, shop_type)
		}
		return *result.OkWithData(ans)
	}
	ans = models.GetShopType()
	for _, j := range ans {
		gredis.Client.SAdd(ctx, utils.SHOP_TYPE_KEY, j)
	}
	return *result.OkWithData(ans)
}
