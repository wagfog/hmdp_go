package service

import (
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
)

type IBLogService interface {
	QueryHotBlog(current int32) result.Result
	QueryBlogById(id int64) result.Result
	LikeBlog(id int64) result.Result
	SaveBLog(blog models.Blog) result.Result
	QueryBlogOfFollow(max int64, offset int32) result.Result
}
