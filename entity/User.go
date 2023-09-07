package entity

import "time"

type User struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Phone      string    `json:"phone"`
	Password   string    `json:"password"`
	NickName   string    `json:"nickName"`
	Icon       string    `json:"icon"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
