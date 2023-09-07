package models

import "time"

type Follow struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID       int64     `json:"userId"`
	FollowUserID int64     `json:"followUserid"`
	CreateTime   time.Time `json:"createTime"`
}
