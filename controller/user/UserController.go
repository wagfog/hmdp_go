package user

import (
	"fmt"
	"log"
	"net/http"

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

func SendCode(c *gin.Context) {
	phone := c.Query("phone")
	if utils.IsPhoneInvalid(phone) {
		c.JSON(http.StatusBadRequest, result.Fail("error phone number"))
		return
	}
	UserService = impl.NewUserService()
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
	UserService = impl.NewUserService()
	ans := UserService.Login(loginFOrmDTO, *sessions.NewSession(nil, ""))
	c.JSON(http.StatusOK, ans)
}
