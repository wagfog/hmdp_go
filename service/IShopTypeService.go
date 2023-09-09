package service

import (
	"github.com/wagfog/hmdp_go/dto/result"
)

type IShopTypeService interface {
	QueryAllList() result.Result
}
