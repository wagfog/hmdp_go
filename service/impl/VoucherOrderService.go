package impl

import (
	"fmt"

	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
)

type VoucherOrderService struct {
}

func (v *VoucherOrderService) SeckillVoucher(voucherId int) result.Result {
	return *result.Ok()
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
