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
