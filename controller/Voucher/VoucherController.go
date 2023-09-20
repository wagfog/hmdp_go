package voucher

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/service"
	"github.com/wagfog/hmdp_go/service/impl"
)

var voucherService service.IVoucherService

func init() {
	voucherService = impl.NewVoucherService()
}

/**
* 新增普通券
* @param voucher 优惠券信息
*@return 优惠券id
 */

func AddVoucher(c *gin.Context) {

}

/**
* 新增秒杀券
* @param voucher 优惠券信息，包含秒杀信息
* @return 优惠券id
 */

func AddSeckillVoucher(c *gin.Context) {
	var voucher models.Voucher
	err := c.ShouldBindJSON(&voucher)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, result.Fail(err.Error()))
		return
	}
	voucherService.AddSeckillVoucher(voucher)
	c.JSON(http.StatusOK, result.Ok())
}

/**
* 查询店铺的优惠券列表
* @param shopId 店铺id
* @return 优惠券列表
 */

func QueryVoucherOfShop(c *gin.Context) {
	sid := c.Param("shopId")
	id, _ := strconv.Atoi(sid)
	res := voucherService.QueryVoucherOfShop(id)
	c.JSON(http.StatusOK, res)
}
