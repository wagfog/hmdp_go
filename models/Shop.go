package models

import "time"

type Shop struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name"`
	TypeID     int64     `json:"typeId"`
	Images     string    `json:"images"`
	Area       string    `json:"area"`
	Address    string    `json:"address"`
	X          float64   `json:"x"`
	Y          float64   `json:"y"`
	AvgPrice   int64     `json:"avgPrice"`
	Sold       int       `json:"sold"`
	Comments   int       `json:"comments"`
	Score      int       `json:"score"`
	OpenHours  string    `json:"opneHours"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Distance   float64   `gorm:"-"`
}
