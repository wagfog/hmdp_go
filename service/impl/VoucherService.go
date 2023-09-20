package impl

import (
	"context"
	"strconv"

	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/utils"
)

type VoucherService struct {
}

func NewVoucherService() *VoucherService {
	return &VoucherService{}
}

func (voucherService *VoucherService) QueryVoucherOfShop(shopId int) result.Result {
	vouchers := models.QueryVoucherOfShop(shopId)
	return *result.OkWithData(vouchers)
}
func (voucherService *VoucherService) AddSeckillVoucher(voucher models.Voucher) result.Result {
	// 保存优惠券
	// models.Db.Begin()
	models.SaveVoucher(&voucher)
	// 保存秒杀信息
	seckillVoucher := models.SeckillVoucher{
		VoucherID: voucher.ID,
		Stock:     voucher.Stock,
		BeginTime: voucher.BeginTime,
		EndTime:   voucher.EndTime,
	}
	models.SaveSeckillVoucher(&seckillVoucher)
	key := utils.SECKILL_STOCK_KEY + strconv.Itoa(int(voucher.ID))
	gredis.Client.Set(context.Background(), key, voucher.Stock, 0)
	return *result.Ok()
}
func (voucherService *VoucherService) SeckillVoucher(voucherId int) result.Result {
	return *result.Fail("not finish")
}
