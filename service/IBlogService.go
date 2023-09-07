package service

import (
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
)

type IBLogService interface {
	queryHotBlog(current int32) result.Result
	queryBlogById(id int64) result.Result
	likeBlog(id int64) result.Result
	saveBLog(blog models.Blog) result.Result
	queryBlogOfFollow(max int64, offset int32) result.Result
}
