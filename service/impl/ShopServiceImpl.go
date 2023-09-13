package impl

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/utils"
)

type ShopServiceImpl struct {
}

func NewShopService() *ShopServiceImpl {
	return &ShopServiceImpl{}
}

func (shopServiceImpl *ShopServiceImpl) QueryById(id int64) result.Result {
	return *result.Fail("not finish")
}

func (shopServiceImpl *ShopServiceImpl) Update(shop models.Shop) result.Result {
	return *result.Fail("not finish")
}

func getEnoughSlice(id []string) []string {
	idx := -1
	for i, s := range id {
		if s == "" {
			break
		}
		idx = i
	}
	if idx == -1 {
		return make([]string, 0)
	}
	return id[0:idx]
}

func (shopServiceImpl *ShopServiceImpl) QueryShopByType(typeId int, current int, x float64, y float64) result.Result {
	if x == 0 || y == 0 {
		shops := models.QueryShopsByType(typeId, current)
		return *result.OkWithData(shops)
	}
	//基于位置的服务（Location Based Services，LBS）
	from := (current - 1) * utils.DEFAULT_PAGE_SIZE
	end := current * utils.DEFAULT_PAGE_SIZE
	Geokey := utils.SHOP_GEO_KEY + strconv.Itoa(typeId)
	resRadiu, err := gredis.Client.GeoRadius(context.Background(), Geokey, x, y, &redis.GeoRadiusQuery{
		Radius:      5000,  //radius表示范围距离，
		Unit:        "m",   //距离单位是 m|km|ft|mi
		WithCoord:   true,  //传入WITHCOORD参数，则返回结果会带上匹配位置的经纬度
		WithDist:    true,  //传入WITHDIST参数，则返回结果会带上匹配位置与给定地理位置的距离。
		WithGeoHash: true,  //传入WITHHASH参数，则返回结果会带上匹配位置的hash值。
		Count:       end,   //入COUNT参数，可以返回指定数量的结果。
		Sort:        "ASC", //默认结果是未排序的，传入ASC为从近到远排序，传入DESC为从远到近排序。
	}).Result()
	if err != nil {
		fmt.Println(err)
	}
	if resRadiu == nil {
		tmp := make([]redis.GeoLocation, 0)
		return *result.OkWithData(tmp)
	}

	if len(resRadiu) < from {
		//没有下一页了，结束
		return *result.Ok()
	}
	//截取 from ~ end的部分
	shopId := make([]string, utils.DEFAULT_PAGE_SIZE)
	shopDistance := make([]float64, utils.DEFAULT_PAGE_SIZE)
	for i, location := range resRadiu {
		if i < from {
			continue
		}
		idx := i - from
		fmt.Println(location.Name, location.GeoHash, location.Dist)
		shopId[idx] = location.Name
		shopDistance[idx] = location.Dist
	}
	realShopId := getEnoughSlice(shopId)
	if len(realShopId) == 0 {
		return *result.Ok()
	}
	idstr := strings.Join(realShopId, ",")

	shops := models.QueryShopsByIds(idstr, shopId)
	fmt.Println(idstr)
	for i, shop := range shops {
		if i < from {
			continue
		}
		idx := i - from
		shop.Distance = shopDistance[idx]
	}
	return *result.OkWithData(shops)
}
