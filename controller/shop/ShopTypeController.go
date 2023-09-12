package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/service"
	"github.com/wagfog/hmdp_go/service/impl"
)

var shopTypeService service.IShopTypeService

func ShopTypeController(c *gin.Context) {
	shopTypeService = impl.NewShopTypeService()
	res := shopTypeService.QueryAllList()
	if !res.Success {
		c.JSON(http.StatusBadRequest, result.NewResultFail("get shop type error!"))
	}
	c.JSON(http.StatusOK, result.NewResultOk(res.Data))
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
	x := c.Query("x")
	y := c.Query("y")

}
