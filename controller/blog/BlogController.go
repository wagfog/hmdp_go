package blog

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/service"
	"github.com/wagfog/hmdp_go/service/impl"
)

var blogService service.IBLogService

func InitBlogService() {
	blogService = impl.NewBlogServiceImpl()
}

func QueryHotBlogController(c *gin.Context) {
	current := c.Query("current")
	cur, _ := strconv.Atoi(current)
	res := blogService.QueryHotBlog(int32(cur))
	c.JSON(http.StatusOK, res)
}

func QueryMyBlog(c *gin.Context) {
	cookie, err := c.Cookie("user_cookie")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadGateway, result.Fail(err.Error()))
		return
	}
	session := sessions.Default(c)
	phone := session.Get(cookie)
	u := models.GetUserByPhone(phone.(string))
	cuurent := c.DefaultQuery("current", "1")
	cur, _ := strconv.Atoi(cuurent)
	res := blogService.QueryBlogByUserId(u.ID, cur)
	c.JSON(http.StatusOK, res)
}

func QueryBlogByUserId(c *gin.Context) {
	current := c.DefaultQuery("current", "1")
	sid := c.Query("id")
	cur, _ := strconv.Atoi(current)
	id, _ := strconv.Atoi(sid)
	if id <= 0 {
		c.JSON(http.StatusBadRequest, result.Fail("user_id error"))
	}
	res := blogService.QueryBlogByUserId(int64(id), cur)
	c.JSON(http.StatusOK, res)
}

func QueryBlogById(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	res := blogService.QueryBlogById(id)
	c.JSON(http.StatusOK, res)
}
