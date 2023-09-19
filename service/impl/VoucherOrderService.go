package impl

import (
	"context"
	"fmt"
	"strconv"

	lua "github.com/wagfog/hmdp_go/config/Lua"
	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/utils"
)

type VoucherOrderService struct {
}

var seckillHash string

func CreateSript() {
	//返回的脚本会产生一个sha1哈希值,下次用的时候可以直接使用这个值
	seckillHash, err := gredis.Client.ScriptLoad(context.Background(), lua.Seckkill).Result()
	if err != nil {
		panic(err)
	}
}

func NewVoucherService(vs *VoucherOrderService) bool {
	if vs != nil {
		return false
	}
	vs = &VoucherOrderService{}
	return true
}

func (v *VoucherOrderService) SeckillVoucher(voucherId int, phone string) result.Result {
	//获取用户
	u := models.GetUserByPhone(phone)
	//获取订单
	orderID := utils.NextId("order")
	id := strconv.Itoa(int(u.ID))
	vid := strconv.Itoa(voucherId)
	sorderID := strconv.Itoa(int(orderID))
	n, err := gredis.Client.EvalSha(context.Background(), seckillHash, []string{id, vid, sorderID}).Result()

	if err != nil {
		panic(err)
	}

	if n == 1 {
		return *result.Fail("库存不足")
	} else if n == 2 {
		return *result.Fail("重复下单")
	}
	return *result.OkWithData(orderID)
}

func (v *VoucherOrderService) CreateVoucherOrder(voucherOrder models.VoucherOrder, phone string) result.Result {
	u := models.GetUserByPhone(phone)
	flag := models.CreateVoucherOrder(int(u.ID), voucherOrder)
	if flag {
		return *result.Ok()
	}
	fmt.Println("库存不足！")
	return *result.Fail("库存不足！")
}
