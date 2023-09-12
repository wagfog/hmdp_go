package main

import (
	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/config/setting"
	"github.com/wagfog/hmdp_go/controller"
	"github.com/wagfog/hmdp_go/controller/user"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/utils"
)

func main() {
	setting.Init()
	models.Init()
	gredis.Setup()
	user.Init()
	utils.InitRedistore()
	server := controller.InitRouter()
	server.Run()
}
