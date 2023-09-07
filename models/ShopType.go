package models

import "time"

type ShopType struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name"`
	Icon       string    `json:"icon"`
	Sort       int       `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}
