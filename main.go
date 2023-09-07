package main

import (
	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/config/setting"
	"github.com/wagfog/hmdp_go/controller"
)

func main() {
	setting.Init()
	gredis.Setup()
	server := controller.InitRouter()
	server.Run()
}
