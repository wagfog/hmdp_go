package models

import (
	"time"

	"github.com/wagfog/hmdp_go/utils"
)

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

func QueryBlogHot(current int32, maxPage int32) []Blog {
	var blogs []Blog
	db.Table("tb_blog").Order("liked").Offset((current - 1) * maxPage).Limit(maxPage).Find(&blogs)
	return blogs
}

func QueryBlogByUser(id int64, current int) []Blog {
	var blogs []Blog
	db.Table("tb_blog").Where("user_id = ?", id).Offset((current - 1) * utils.MAX_PAGE_SIZE).Limit(utils.MAX_PAGE_SIZE).Find(&blogs)
	return blogs
}

func QueryBlogById(id int) *Blog {
	var blogs Blog
	db.Table("tb_blog").Where("id = ?", id).First(&blogs)
	return &blogs
}
