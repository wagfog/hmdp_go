package service

import (
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
)

type IBLogService interface {
	QueryHotBlog(current int32) result.Result
	QueryBlogByUserId(id int64, current int) result.Result
	QueryBlogById(id int, phone string) result.Result
	LikeBlog(id int64, phone string) result.Result
	SaveBLog(blog models.Blog) result.Result
	QueryBlogOfFollow(max int64, offset int64, phone string) result.Result
	QueryBlogLike(id int) result.Result
}
