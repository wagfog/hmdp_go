package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "文件上传失败")
		return
	}
	defer file.Close()

	// 获取原始文件名称
	originalFilename := handler.Filename
	// 生成新文件名
}

func CreateNewFileName(originalFilename string) {
	// 获取后缀
}
