package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/service"
	"github.com/wagfog/hmdp_go/service/impl"
	"github.com/wagfog/hmdp_go/utils"
)

var (
	UserService service.IUserService
)

func SendCode(c *gin.Context) {
	phone := c.Query("phone")
	if utils.IsPhoneInvalid(phone) {
		c.JSON(http.StatusBadRequest, result.Fail("error phone number"))
		return
	}
	UserService = impl.NewUserService()
	UserService.SendCode(phone, *sessions.NewSession(nil, ""))
}
