package image_api

import (
	"github.com/gin-gonic/gin"
	"we_a_family/we_a_family/service"
)

type ImageApi struct {
}

// ImageUploadView 上传单个图片
func (ImageApi) ImageUploadView(ctx *gin.Context) {
	service.ImageUploadService(ctx)

}

// ImageDownloadView 下载单个图片(根据照片名字)
func (ImageApi) ImageDownloadView(ctx *gin.Context) {
	service.ImageDownloadService(ctx)

}

// ImageFindAllView 查找目录下所有图片
func (ImageApi) ImageFindAllView(ctx *gin.Context) {
	service.ImageFindAllService(ctx)

}

func (ImageApi) ImageDeleteView(ctx *gin.Context) {
	service.ImageDeleteService(ctx)
}

func (ImageApi) ImageUpdateView(ctx *gin.Context) {
	service.ImageUpdateService(ctx)
}
