package main

import (
	"github.com/wagfog/hmdp_go/controller"
)

func main() {
	server := controller.InitRouter()
	server.Run()
}
