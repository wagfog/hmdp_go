package impl

import (
	"context"
	"strconv"
	"time"

	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
)

type FollowService struct {
}

func NewFollowService() *FollowService {
	return &FollowService{}
}

func (followService *FollowService) IsFollow(followUserid int, phone string) result.Result {
	user := models.GetUserByPhone(phone)
	count := models.IsFollow(user.ID, int64(followUserid))
	return *result.OkWithData(count > 0)
}

func (followService *FollowService) Follow(followUserid int, isFollow bool, phone string) result.Result {
	user := models.GetUserByPhone(phone)
	key := "follows:" + strconv.Itoa(int(user.ID))
	followUseridStr := strconv.Itoa(followUserid)
	//判断到底是关注还是取消关注
	if isFollow {
		//关注：新增数据
		follow := models.Follow{
			UserID:       user.ID,
			FollowUserID: int64(followUserid),
			CreateTime:   time.Now(),
		}
		models.SaveFollow(&follow)
		gredis.Client.SAdd(context.Background(), key, followUseridStr)
	} else {
		//取关，删除
		models.RemoveFollow(user.ID, int64(followUserid))
		gredis.Client.SRem(context.Background(), key, followUseridStr)
	}
	return *result.Ok()
}

func (followService *FollowService) FollowCommons(id int) result.Result {
	return *result.Fail("not finish")
}
