package service

import "github.com/wagfog/hmdp_go/dto/result"

type IFollowService interface {
	IsFollow(followUserid int, phone string) result.Result

	Follow(followUserid int, isFollow bool, phone string) result.Result

	FollowCommons(id int) result.Result
}
