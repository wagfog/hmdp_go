package impl

import (
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

func (BlogServiceImpl *BlogServiceImpl) QueryBlogById(id int) result.Result {
	blog := models.QueryBlogById(id)
	if blog == nil {
		return *result.Fail("笔记不存在!")
	}
	queryBlogUser(blog)
	isBlogLiked(blog)
	return *result.OkWithData(blog)
}

func (BlogServiceImpl *BlogServiceImpl) LikeBlog(id int64) result.Result {
	return *result.Ok()
}
func (BlogServiceImpl *BlogServiceImpl) SaveBLog(blog models.Blog) result.Result {
	return *result.Ok()
}
func (BlogServiceImpl *BlogServiceImpl) QueryBlogOfFollow(max int64, offset int32) result.Result {
	return *result.Ok()
}

func queryBlogUser(blog *models.Blog) {
	userId := blog.ID
	user := models.GetUser(userId)
	blog.Name = user.NickName
	blog.Icon = user.Icon
}

func isBlogLiked(blog *models.Blog) {

}
