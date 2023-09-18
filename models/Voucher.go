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
	Status      int       `json:"status"`
	Stock       int       `json:"-"`
	BeginTime   time.Time `json:"-"`
	EndTime     time.Time `json:"-"`
	CreateTime  time.Time `json:"createTime" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"updateTime" gorm:"column:create_time"`
}

func QueryVoucherOfShop(shopId int) []Voucher {
	var vouchers []Voucher
	db.Table("tb_voucher").Where("shop_id = ?", shopId).Find(&vouchers)
	return vouchers
}

func SaveVoucher(v *Voucher) {
	db.Table("tb_voucher").Create(v)
}
