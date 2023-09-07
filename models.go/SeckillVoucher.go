package models

import "time"

type SeckillVoucher struct {
	VoucherID  int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Stock      int       `json:"stock"`
	CreateTime time.Time `json:"createTime"`
	BeginTime  time.Time `json:"beginTime"`
	EndTime    time.Time `json:"endTime"`
	UpdateTime time.Time `json:"updateTime"`
}
