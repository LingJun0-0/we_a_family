package routers

import (
	"we_a_family/we_a_family/api"
	"we_a_family/we_a_family/service"
)

func (router RouterGroup) TagRouter() {
	tagApi := api.GroupApp.TagApi
	router.Use(service.Auth)
	router.GET("tag/find", tagApi.FindTagView)
	router.GET("tag/findAll", tagApi.FindAllTagView)
	router.POST("tag/create", tagApi.CreateTagView)
	router.PATCH("tag/update", tagApi.UpdateTagView)
	router.DELETE("tag/delete", tagApi.DeleteTagView)

}
