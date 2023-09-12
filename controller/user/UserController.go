package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/wagfog/hmdp_go/dto"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/service"
	"github.com/wagfog/hmdp_go/service/impl"
	"github.com/wagfog/hmdp_go/utils"
)

var (
	UserService service.IUserService //需要统一初始化一下，或者用某种框架解耦哦
)

func Init() {
	UserService = impl.NewUserService()
}

func SendCode(c *gin.Context) {
	phone := c.Query("phone")
	if utils.IsPhoneInvalid(phone) {
		c.JSON(http.StatusBadRequest, result.Fail("error phone number"))
		return
	}
	ans := UserService.SendCode(phone, *sessions.NewSession(nil, ""))
	c.JSON(http.StatusOK, ans)
}

func Login(c *gin.Context) {
	// phone := c.Query("phone")
	// if utils.IsPhoneInvalid(phone) {
	// 	c.JSON(http.StatusBadRequest, result.Fail("error phone number"))
	// }

	// loginFOrmDTO := dto.NewLoginFOrmDTO(c.PostForm("phone"), c.PostForm("code"), c.PostForm("password"))
	var loginFOrmDTO dto.LoginFormDTO2
	err := c.ShouldBindJSON(&loginFOrmDTO)
	if err != nil {
		log.Fatal("error", err)
		c.JSON(http.StatusBadRequest, result.Fail("bad request"))
		return
	}

	fmt.Println(loginFOrmDTO.GetCode(), loginFOrmDTO.GetPhone())
	if loginFOrmDTO.GetCode() == "" {
		c.JSON(http.StatusBadRequest, result.Fail("bad request"))
		return
	}

	ans := UserService.Login(loginFOrmDTO, *sessions.NewSession(nil, ""))
	c.JSON(http.StatusOK, ans)
}

func Me(c *gin.Context) {
	u := dto.UserDTO{
		Id:       4,
		NickName: "user_slxaxy2au9f3tanffaxr",
		Icon:     "",
	}
	fmt.Println("use this me")
	udto := dto.UserDTO{
		Id:       u.Id,
		NickName: u.NickName,
		Icon:     u.Icon,
	}
	c.JSON(http.StatusOK, result.OkWithData(udto))
}

func Info(c *gin.Context) {
	fmt.Println("use Info")
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	if id <= 0 {
		c.JSON(http.StatusOK, result.Ok())
	}
	fmt.Println(sid)
	u := UserService.GetUserById(int64(id))
	udto := dto.UserDTO{
		Id:       u.ID,
		NickName: u.NickName,
		Icon:     u.Icon,
	}
	fmt.Println(udto)
	c.JSON(http.StatusOK, result.OkWithData(udto))
}

func QueryUserById(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	u := UserService.GetUserById(int64(id))
	c.JSON(http.StatusOK, result.OkWithData(u))
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, result.Ok())
}
