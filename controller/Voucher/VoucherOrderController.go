package voucher

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/service"
	"github.com/wagfog/hmdp_go/service/impl"
)

var voucherOrderService service.IVoucherOrderService

func Init() {
	voucherOrderService = impl.NewVoucherOrderService()
	impl.CreateSript()
}
func SeckillVoucher(c *gin.Context) {
	sid := c.Param("id")
	session := sessions.Default(c)
	cookie, err := c.Cookie("user_cookie")
	fmt.Println(cookie)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, result.Fail(err.Error()))
		return
	}
	phone := session.Get(cookie)
	id, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, result.Fail(err.Error()))
		return
	}
	fmt.Println(id)
	res := voucherOrderService.SeckillVoucher(id, phone.(string))
	c.JSON(http.StatusOK, res)
}
