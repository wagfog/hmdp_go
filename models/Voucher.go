package models

import "time"

type Voucher struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ShopID      int64     `json:"shopId" gorm:"column:shop_id"`
	Title       string    `json:"title"`
	SubTitle    string    `json:"subTitle" gorm:"column:sub_title"`
	Rules       string    `json:"rules"`
	PayValue    int64     `json:"payValue" gorm:"column:pay_value"`
	ActualValue int64     `json:"actualValue" gorm:"column:actual_value"`
	Type        int       `json:"type"`
	Status      int       `json:"-"`
	Stock       int       `json:"stock"`
	BeginTime   time.Time `json:"beginTime"`
	EndTime     time.Time `json:"endTime"`
	CreateTime  time.Time `json:"-" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"-" gorm:"column:create_time"`
}

func QueryVoucherOfShop(shopId int) []Voucher {
	var vouchers []Voucher
	db.Table("tb_voucher").
		Select("tb_voucher.*,tb_seckill_voucher.stock, tb_seckill_voucher.begin_time, tb_seckill_voucher.end_time").
		Joins("LEFT JOIN tb_seckill_voucher tb_seckill_voucher ON tb_voucher.id = tb_seckill_voucher.voucher_id").
		Where("tb_voucher.shop_id = ? AND tb_voucher.status = ?", shopId, 1).
		Find(&vouchers)
	return vouchers
}

func SaveVoucher(v *Voucher) {
	db.Table("tb_voucher").Create(v)
}
