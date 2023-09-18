package service

import (
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
)

type IVoucherService interface {
	QueryVoucherOfShop(shopId int) result.Result
	AddSeckillVoucher(voucher models.Voucher) result.Result
	SeckillVoucher(voucherId int) result.Result
}
