package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/controller/user"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	userGroup := r.Group("/user")
	userGroup.POST("/code", user.Login)

	return r
}
