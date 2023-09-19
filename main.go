package main

import (
	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/config/rabbitmq"
	"github.com/wagfog/hmdp_go/config/setting"
	"github.com/wagfog/hmdp_go/controller"
	"github.com/wagfog/hmdp_go/controller/blog"
	"github.com/wagfog/hmdp_go/controller/follow"
	"github.com/wagfog/hmdp_go/controller/shop"
	"github.com/wagfog/hmdp_go/controller/user"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/utils"
)

func main() {
	setting.Init()
	rabbitmq.Init_Rabbitmq()
	defer rabbitmq.Mq_Conn.Close()
	models.Init()
	gredis.Setup()
	user.Init()
	blog.InitBlogService()
	utils.InitRedistore()
	shop.Init()
	follow.Init()
	server := controller.InitRouter()
	server.Run()
}
