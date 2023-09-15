package impl

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/dto"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/utils"
)

type BlogServiceImpl struct {
}

func NewBlogServiceImpl() *BlogServiceImpl {
	return &BlogServiceImpl{}
}

func (BlogServiceImpl *BlogServiceImpl) QueryHotBlog(current int32) result.Result {
	blogs := models.QueryBlogHot(current, utils.MAX_PAGE_SIZE)

	for _, blog := range blogs {
		queryBlogUser(&blog)
	}
	return *result.OkWithData(blogs)
}
func (BlogServiceImpl *BlogServiceImpl) QueryBlogByUserId(id int64, current int) result.Result {
	blogs := models.QueryBlogByUser(id, current)
	return *result.OkWithData(blogs)
}

func (BlogServiceImpl *BlogServiceImpl) QueryBlogById(id int, phone string) result.Result {
	u := models.GetUserByPhone(phone)
	blog := models.QueryBlogById(id)
	if blog == nil {
		return *result.Fail("笔记不存在!")
	}
	queryBlogUser(blog)
	isBlogLiked(blog, u)
	return *result.OkWithData(blog)
}

func (BlogServiceImpl *BlogServiceImpl) LikeBlog(id int64, phone string) result.Result {
	user := models.GetUserByPhone(phone)
	key := utils.BLOG_LIKED_KEY + strconv.Itoa(int(id))
	userId := strconv.Itoa(int(user.ID))
	_, err := gredis.Client.ZScore(context.Background(), key, userId).Result()
	currntTime := time.Now().UnixNano() / 1000000
	if err == redis.Nil {
		fmt.Println("如果未点赞，可以点赞")
		//如果未点赞，可以点赞
		//数据库点赞数量+1
		models.UpdateBlogLikePlus(int(id))
		_, zaddErr := gredis.Client.ZAdd(context.Background(), key, &redis.Z{
			Score:  float64(currntTime),
			Member: userId,
		}).Result()
		if zaddErr != nil {
			fmt.Println(err)
		}
	} else {
		//如果已经点赞，取消点赞
		//数据库点赞-1
		fmt.Println("如果已经点赞，取消点赞")
		models.UpdateBlogLikeSub(int(id))
		gredis.Client.ZRem(context.Background(), key, userId)
	}
	return *result.Ok()
}
func (BlogServiceImpl *BlogServiceImpl) SaveBLog(blog models.Blog) result.Result {
	return *result.Ok()
}
func (BlogServiceImpl *BlogServiceImpl) QueryBlogOfFollow(max int64, offset int64, phone string) result.Result {
	//1.获取当前用户
	user := models.GetUserByPhone(phone)

	//2.查询收件箱 ZREVRANGBYSCORE key Max Min LIMIT offset count
	key := utils.FEED_KEY + strconv.Itoa(int(user.ID))
	res, err := gredis.Client.ZRevRangeByScoreWithScores(context.Background(), key, &redis.ZRangeBy{
		Min:    "0",
		Max:    strconv.Itoa(int(max)),
		Offset: offset,
		Count:  2,
	}).Result()

	// 处理错误
	if err != nil {
		return *result.Fail(err.Error())
	}
	if res == nil || len(res) == 0 {
		return *result.Ok()
	}
	var (
		minTime int64 = 0
		os      int   = 1
	)
	ids := make([]string, len(res))
	fmt.Println(res)
	for i, tuple := range res {
		id := (tuple.Member).(string)
		ids[i] = id
		val := int64(tuple.Score)
		if val == minTime {
			os++
		} else {
			minTime = val
			os = 1
		}
	}
	strIds := strings.Join(ids, ",")
	blogs := models.QueryBlogByIds(ids, strIds)

	for _, blog := range blogs {
		queryBlogUser(&blog)
		isBlogLiked(&blog, user)
	}
	inter := make([]interface{}, len(blogs))
	for i, blog := range blogs {
		inter[i] = blog
	}
	scrollRes := dto.ScrollResult{
		List:    inter,
		MinTime: minTime,
		Offset:  os,
	}
	return *result.OkWithData(scrollRes)
}

func (BlogServiceImpl *BlogServiceImpl) QueryBlogLike(id int) result.Result {
	//1.查询top5点赞用户 zrange key 0 4
	key := utils.BLOG_LIKED_KEY + strconv.Itoa(id)
	TopUid, err := gredis.Client.ZRange(context.Background(), key, 0, 4).Result()
	if err != nil {
		fmt.Println(err)
		return *result.Fail(err.Error())
	}
	if TopUid == nil {
		return *result.Ok()
	}
	idStr := strings.Join(TopUid, ",")
	users := models.QueryUsersByIds(idStr, TopUid)
	return *result.OkWithData(users)
}

func queryBlogUser(blog *models.Blog) {
	userId := blog.UserID
	user := models.GetUser(userId)
	blog.Name = user.NickName
	blog.Icon = user.Icon
}

func isBlogLiked(blog *models.Blog, u *models.User) {
	if u == nil {
		return
	}
	uID := strconv.Itoa(int(u.ID))
	key := utils.BLOG_LIKED_KEY + strconv.Itoa(int(blog.ID))
	_, err := gredis.Client.ZScore(context.Background(), key, uID).Result()
	if err == redis.Nil {
		blog.IsLike = false
	}
	blog.IsLike = true
}
