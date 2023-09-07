package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/utils"
)

func Login(c *gin.Context) {
	phone := c.Query("phone")
	if utils.IsPhoneInvalid(phone) {
		c.JSON(http.StatusBadRequest, result.Fail("error phone number"))
		return
	}
	c.JSON(http.StatusOK, result.OkWithData("not finish"))
}
