package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

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
	shop, err := queruWithMutex(id)
	if shop == nil {
		return *result.Fail("不存在该商铺信息！")
	}
	if err != nil {
		return *result.Fail(err.Error())
	}
	return *result.OkWithData(*shop)
}

// 互斥锁解决缓存穿透的方法
func queruWithMutex(id int64) (*models.Shop, error) {
	key := utils.CACHE_SHOP_KEY + strconv.Itoa(int(id))
	fmt.Println(key)
	shopJson, err := gredis.Client.Get(context.Background(), key).Result()
	if err == nil {
		var shop models.Shop
		fmt.Println(shopJson)
		er := json.Unmarshal([]byte(shopJson), &shop)
		fmt.Println(shop)
		if er != nil {
			fmt.Println(er)
			return nil, er
		}
		return &shop, nil
	}
	fmt.Println("not get any")
	//获取互斥锁
	LockKey := utils.LOCK_SHOP_KEY + strconv.Itoa(int(id))
	isLock := tryLock(LockKey)
	//判断是否获取
	if !isLock {
		//失败，休眠并重试
		time.Sleep(5 * time.Microsecond)
		return queruWithMutex(id)
	}
	//成功，根据id查询数据库
	shop := models.QueryShopByid(id)
	//不存在，返回错误,并将空值写入redis
	if shop == nil {
		gredis.Client.SetEX(context.Background(), key, utils.CACHE_NULL_TTL, time.Minute)
		return nil, nil
	}
	//存在，写入redis
	bJson, _ := json.Marshal(*shop)
	JsonStr := string(bJson)
	gredis.Client.Set(context.Background(), key, JsonStr, time.Minute)
	unLock(LockKey)
	return shop, nil
}

func tryLock(key string) bool {
	flag, _ := gredis.Client.SetNX(context.Background(), key, "1", 10*time.Second).Result()
	return flag
}

func unLock(key string) {
	gredis.Client.Del(context.Background(), key)
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
