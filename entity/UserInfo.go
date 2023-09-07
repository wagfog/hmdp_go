package entity

import "time"

type UserInfo struct {
	UserID     int64     `json:"userId" gorm:"primaryKey;autoIncrement"`
	City       string    `json:"city"`
	Introduce  string    `json:"introduce"`
	Fans       int       `json:"fans"`
	Followee   int       `json:"followee"`
	Gender     bool      `json:"gender"`
	Birthday   time.Time `json:"birthday"`
	Credits    int       `json:"credits"`
	Level      bool      `json:"level"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
