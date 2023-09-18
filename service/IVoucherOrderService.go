package service

import (
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
)

type IVoucherOrderService interface {
	SeckillVoucher(voucherId int) result.Result

	CreateVoucherOrder(voucherOrder models.VoucherOrder, phone string)
}
