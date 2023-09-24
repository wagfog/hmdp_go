package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	voucher "github.com/wagfog/hmdp_go/controller/Voucher"
	"github.com/wagfog/hmdp_go/controller/blog"
	"github.com/wagfog/hmdp_go/controller/follow"
	"github.com/wagfog/hmdp_go/controller/shop"
	"github.com/wagfog/hmdp_go/controller/user"
	"github.com/wagfog/hmdp_go/utils"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
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

	shopGroup := r.Group("/shop")
	shopGroup.GET("/of/type", shop.QueryShopByType)
	shopGroup.GET("/:id", shop.QueryShopById)

	voucherGroup := r.Group("/voucher")
	voucherGroup.GET("/list/:shopId", voucher.QueryVoucherOfShop)
	voucherGroup.POST("/seckill", voucher.AddSeckillVoucher)

	voucherOrderGroup := r.Group("/voucher-order")
	voucherOrderGroup.Use(CookieAuth())
	voucherOrderGroup.POST("/seckill/:id", voucher.SeckillVoucher)

	blogGroup := r.Group("/blog")
	blogGroup.Use(CookieAuth())
	// blogGroup.GET("/hot", blog.QueryHotBlogController)
	r.GET("/blog/hot", blog.QueryHotBlogController)
	blogGroup.GET("/of/me", blog.QueryMyBlog)
	blogGroup.GET("/:id", blog.QueryBlogById)
	blogGroup.GET("/of/user", blog.QueryBlogByUserId)
	blogGroup.GET("/of/follow", blog.QueryBlogOfFollow)
	blogGroup.GET("/likes/:id", blog.QueryBlogLike)
	blogGroup.PUT("/like/:id", blog.LikeBlog)

	followGroup := r.Group("/follow")
	followGroup.Use(CookieAuth())
	followGroup.PUT("/:id/:isFollow", follow.Follow)
	followGroup.GET("/or/not/:id", follow.IsFollow)

	uploadGroup := r.Group("/upload")
	uploadGroup.Use(CookieAuth())
	uploadGroup.POST("/blog")

	return r
}
