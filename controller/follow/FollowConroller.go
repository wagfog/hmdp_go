package follow

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/service"
	"github.com/wagfog/hmdp_go/service/impl"
)

var FollowService service.IFollowService

func Init() {
	FollowService = impl.NewFollowService()
}

func Follow(c *gin.Context) {
	sid := c.Param("id")
	isFollow := c.Param("isFollow")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Fail("查询id错误"))
	}
	FollowIs, _ := strconv.ParseBool(isFollow)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Fail("是否关注参数错误"))
	}
	str, err := c.Cookie("user_cookie")
	session := sessions.Default(c)
	userPhone := session.Get(str)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Fail(err.Error()))
	}
	res := FollowService.Follow(id, FollowIs, userPhone.(string))
	c.JSON(http.StatusOK, res)
}

func IsFollow(c *gin.Context) {
	sfollowUid := c.Param("id")
	sesssion := sessions.Default(c)
	cookieId, er := c.Cookie("user_cookie")
	if er != nil {
		c.JSON(http.StatusBadRequest, result.Fail(er.Error()))
	}
	followUid, err := strconv.Atoi(sfollowUid)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Fail(err.Error()))
	}
	phone := sesssion.Get(cookieId)
	res := FollowService.IsFollow(followUid, phone.(string))
	c.JSON(http.StatusOK, res)
}
