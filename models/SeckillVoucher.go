package models

import "time"

type SeckillVoucher struct {
	VoucherID  int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Stock      int       `json:"stock"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`
	BeginTime  time.Time `json:"beginTime" gorm:"column:begin_time"`
	EndTime    time.Time `json:"endTime" gorm:"column:end_time"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"`
}

func SaveSeckillVoucher(v *SeckillVoucher) {
	db.Table("tb_seckill_voucher").Create(v)
}
