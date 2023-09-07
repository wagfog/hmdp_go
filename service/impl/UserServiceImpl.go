package impl

import (
	"github.com/gorilla/sessions"
	"github.com/wagfog/hmdp_go/dto"
	"github.com/wagfog/hmdp_go/dto/result"
)

type UserServiceImpl struct {
}

func (serServiceImpl *UserServiceImpl) SendCode(phone string, sess sessions.Session) result.Result {
	return result.Ok()
}
func (serServiceImpl *UserServiceImpl) Login(loginForm dto.LoginFormDTO, sess sessions.Session) result.Result {
	return result.Ok()
}
func (serServiceImpl *UserServiceImpl) Sign() result.Result {
	return result.Ok()
}

func (serServiceImpl *UserServiceImpl) SignCount() result.Result {
	return result.Ok()
}
