package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type VoucherOrder struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     int64     `json:"userId"`
	VoucherID  int64     `json:"voucherId"`
	PayType    int       `json:"payType"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	PayTime    time.Time `json:"payTime"`
	RefundTime time.Time `json:"refundTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func CreateVoucherOrder(uid int, v VoucherOrder) bool {
	tx := db.Begin()
	defer tx.Commit()
	var count int
	tx.Table("tb_voucher_order").Where("user_id = ? and voucher_id = ?", uid, v.ID).Count(&count)
	if count > 0 {
		fmt.Println("用户已经购买过一次了")
		return false
	}

	//扣减库存
	res := tx.Table("tb_voucher_order").Where("voucher_id = ?", v.ID).
		Where("stock > 0").
		Update("stock", gorm.Expr("stock - ?", 1))

	if res.Error != nil || res.RowsAffected == 0 { // 更新失败或库存不足
		return false
	}

	tx.Table("tb_voucher_order").Create(v)
	return true
}
