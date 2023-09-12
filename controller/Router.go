package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/wagfog/hmdp_go/controller/blog"
	"github.com/wagfog/hmdp_go/controller/shop"
	"github.com/wagfog/hmdp_go/controller/user"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	userGroup := r.Group("/user")
	userGroup.POST("/code", user.SendCode)
	userGroup.POST("/login", user.Login)
	userGroup.GET("/me", user.Me)
	userGroup.GET("/info/:id", user.Info)
	userGroup.GET("/:id", user.QueryUserById)
	userGroup.POST("/logout", user.Logout)

	shopTypeGroup := r.Group("/shop-type")
	shopTypeGroup.GET("/list", shop.ShopTypeController)

	blogGroup := r.Group("/blog")
	blogGroup.GET("/hot", blog.QueryHotBlogController)
	return r
}
