package service

import (
	"github.com/gorilla/sessions"
	"github.com/wagfog/hmdp_go/dto"
	"github.com/wagfog/hmdp_go/dto/result"
)

type IUserService interface {
	SendCode(phone string, sess sessions.Session) result.Result
	Login(loginForm dto.LoginFormDTO, sess sessions.Session) result.Result
	Sign() result.Result
	SignCount() result.Result
}
