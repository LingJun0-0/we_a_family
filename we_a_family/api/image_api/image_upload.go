package image_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"we_a_family/we_a_family/utils"
)

// ImageUploadView 上传单个图片，返沪以图片的URL
func (ImageApi) ImageUploadView(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	fmt.Println(fileHeader.Header)
	fmt.Println(fileHeader.Size)
	fmt.Println(fileHeader.Filename)
}
