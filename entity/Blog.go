package entity

import "time"

type Blog struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ShopID     int64     `json:"shopId"`
	UserID     int64     `json:"userId"`
	Icon       string    `json:"-"`
	Name       string    `json:"-"`
	IsLike     bool      `json:"-"`
	Title      string    `json:"title"`
	Images     string    `json:"images"`
	Content    string    `json:"content"`
	Liked      int       `json:"liked"`
	Comments   int       `json:"comments"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
