package routers

import (
	"we_a_family/we_a_family/api"
	"we_a_family/we_a_family/service"
)

func (router RouterGroup) ImageRouter() {
	imagesApi := api.GroupApp.ImagesApi
	router.Use(service.Auth)
	router.GET("images/findAll", imagesApi.ImageFindAllView)
	router.POST("images/upload", imagesApi.ImageUploadView)
	router.GET("images/download/:id", imagesApi.ImageDownloadView)
	router.PATCH("images/update", imagesApi.ImageUpdateView)
	router.GET("images/download", imagesApi.ImageDownloadView)
	router.DELETE("images/delete/:id", imagesApi.ImageDeleteView)
}
