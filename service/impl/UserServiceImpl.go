package impl

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/dto"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/utils"
)

type UserServiceImpl struct {
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (userServiceImpl *UserServiceImpl) SendCode(phone string, sess sessions.Session) result.Result {
	if utils.IsPhoneInvalid(phone) {
		return *result.Fail("手机号格式错误!")
	}

	rand.Seed(time.Now().UnixNano())

	randomNUm := 0
	for randomNUm < 100000 || randomNUm > 999999 {
		randomNUm = rand.Intn(1000000)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	code := strconv.Itoa(randomNUm)
	gredis.Client.Set(ctx, "login:code:"+phone, code, 2*time.Minute)
	fmt.Println("the code is", code)
	// st, _ := utils.GenerateToken("phone")
	return *result.Ok()
}
func (userServiceImpl *UserServiceImpl) Login(loginForm dto.LoginFormDTO2, sess sessions.Session) result.Result {
	phone := loginForm.GetPhone()
	if utils.IsPhoneInvalid(phone) {
		return *result.Fail("phone error!")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cacheCode, err := gredis.Client.Get(ctx, utils.LOGIN_CODE_KEY+phone).Result()
	if err != nil {
		fmt.Println("redis get cacheCode from redis error!", err)
	}
	code := loginForm.GetCode()

	if cacheCode == "" || code != cacheCode {
		return *result.Fail("cacheCode error or cacheCode is empty")
	}

	user := models.GetUserByPhone(phone)
	if user == nil {
		user, err = models.CreateUser(phone)
		if err != nil || user == nil {
			return *result.Fail("create user Fail" + err.Error())
		}
	}
	// jwtS, err := utils.GenerateToken(phone)
	// if err != nil {
	// 	fmt.Println("jwt error")
	// 	return *result.Fail("get token error")
	// }
	return *result.OkWithData(user)

}
func (userServiceImpl *UserServiceImpl) Sign() result.Result {
	return *result.Ok()
}

func (userServiceImpl *UserServiceImpl) SignCount() result.Result {
	return *result.Ok()
}

func (UserServiceImpl *UserServiceImpl) GetUserById(id int64) models.User {
	return *models.GetUserById(id)
}
