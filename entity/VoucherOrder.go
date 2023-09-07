package entity

import "time"

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
