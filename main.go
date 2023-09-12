package main

import (
	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/config/setting"
	"github.com/wagfog/hmdp_go/controller"
	"github.com/wagfog/hmdp_go/controller/user"
	"github.com/wagfog/hmdp_go/models"
)

func main() {
	setting.Init()
	models.Init()
	gredis.Setup()
	user.Init()
	server := controller.InitRouter()
	server.Run()
}
