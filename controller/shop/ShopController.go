package shop

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/service"
	"github.com/wagfog/hmdp_go/service/impl"
)

var shopService service.IShopService

func Init() {
	shopService = impl.NewShopService()
}

/**
* 根据商铺类型分页查询商铺信息
 * @param typeId 商铺类型
 * @param current 页码
 * @return 商铺列表
*/
func QueryShopByType(c *gin.Context) {
	typeId := c.Query("typeId")
	cur := c.DefaultQuery("current", "1")
	x := c.DefaultQuery("x", "0")
	y := c.DefaultQuery("y", "0")

	tid, _ := strconv.Atoi(typeId)
	current, _ := strconv.Atoi(cur)
	X, _ := strconv.ParseFloat(x, 64)
	Y, _ := strconv.ParseFloat(y, 64)
	res := shopService.QueryShopByType(tid, current, X, Y)
	c.JSON(http.StatusOK, res)
}
