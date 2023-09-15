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
	session := sessions.Default(c)
	cookie, _ := c.Cookie("user_cookie")
	phone := session.Get(cookie)
	res := blogService.QueryBlogById(id, phone.(string))
	c.JSON(http.StatusOK, res)
}

func QueryBlogOfFollow(c *gin.Context) {
	smax := c.Query("lastId")
	sOffset := c.DefaultQuery("offset", "0")
	max, err := strconv.Atoi(smax)
	if err != nil {
		c.JSON(http.StatusBadGateway, result.Fail(err.Error()))
		return
	}
	offset, err2 := strconv.Atoi(sOffset)
	if err2 != nil {
		c.JSON(http.StatusBadGateway, result.Fail(err2.Error()))
		return
	}
	session := sessions.Default(c)
	cookie, _ := c.Cookie("user_cookie")
	phone := session.Get(cookie)
	res := blogService.QueryBlogOfFollow(int64(max), int64(offset), phone.(string))
	c.JSON(http.StatusOK, res)
}

func QueryBlogLike(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(http.StatusBadGateway, result.Fail(err.Error()))
		return
	}
	res := blogService.QueryBlogLike(id)
	c.JSON(http.StatusOK, res)
}

func LikeBlog(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	session := sessions.Default(c)
	cookieId, _ := c.Cookie("user_cookie")
	if err != nil {
		c.JSON(http.StatusBadGateway, result.Fail(err.Error()))
		return
	}
	phone := session.Get(cookieId)
	res := blogService.LikeBlog(int64(id), phone.(string))
	c.JSON(http.StatusOK, res)
}
