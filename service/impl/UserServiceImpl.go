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
	return *result.Ok()
}
func (userServiceImpl *UserServiceImpl) Login(loginForm dto.LoginFormDTO, sess sessions.Session) result.Result {
	return *result.Ok()
}
func (userServiceImpl *UserServiceImpl) Sign() result.Result {
	return *result.Ok()
}

func (userServiceImpl *UserServiceImpl) SignCount() result.Result {
	return *result.Ok()
}
