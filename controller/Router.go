package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/wagfog/hmdp_go/controller/blog"
	"github.com/wagfog/hmdp_go/controller/shop"
	"github.com/wagfog/hmdp_go/controller/user"
	"github.com/wagfog/hmdp_go/utils"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(sessions.Sessions("mysession", utils.Redistore))
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
	blogGroup.GET("/of/me", blog.QueryMyBlog)
	blogGroup.GET("/:id", blog.QueryBlogById)
	blogGroup.GET("/of/user", blog.QueryBlogByUserId)
	return r
}
