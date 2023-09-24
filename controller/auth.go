package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/dto/result"
)

func CookieAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("user_cookie")
		if err != nil || cookie == "" {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, result.Fail(err.Error()))
			ctx.Abort()
			return
		}
		session := sessions.Default(ctx)
		phone := session.Get(cookie)
		if phone == nil {
			ctx.JSON(http.StatusBadRequest, result.Fail(err.Error()))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
