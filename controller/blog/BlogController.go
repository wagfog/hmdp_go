package blog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/service"
	"github.com/wagfog/hmdp_go/service/impl"
)

var blogService service.IBLogService

func QueryHotBlogController(c *gin.Context) {
	blogService = impl.NewBlogServiceImpl()
	current := c.Query("current")
	cur, _ := strconv.Atoi(current)
	res := blogService.QueryHotBlog(int32(cur))
	c.JSON(http.StatusOK, res)
}
