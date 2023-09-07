package models

import "time"

type BlogComment struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	BlogID     int64     `json:"blogId"`
	UserID     int64     `json:"userId"`
	ParentID   int64     `json:"parentId"`
	AnswerID   int64     `json:"answerId"`
	Content    string    `json:"content"`
	Liked      int       `json:"liked"`
	Status     bool      `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
