package entity

import "time"

type Voucher struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ShopID      int64     `json:"shopId"`
	Title       string    `json:"title"`
	SubTitle    string    `json:"subTitle"`
	Rules       string    `json:"rules"`
	PayValue    int64     `json:"payValue"`
	ActualValue int64     `json:"actualValue"`
	Type        int       `json:"type"`
	Status      int       `json:"status"`
	Stock       int       `json:"-"`
	BeginTime   time.Time `json:"-"`
	EndTime     time.Time `json:"-"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}
