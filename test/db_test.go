package test

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/wagfog/hmdp_go/config/setting"
	"github.com/wagfog/hmdp_go/models"
)

func TestDb(t *testing.T) {
	setting.Init()
	models.Init()
	err := models.Db.Table("tb_blog").Where("id = ?", 6).Update("liked", gorm.Expr("liked - ?", 1)).Error
	fmt.Println(err)
	var blog models.Blog
	models.Db.Table("tb_blog").Where("id = ?", 6).First(&blog)
	fmt.Println(blog.Liked)
}
