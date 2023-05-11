package routers

import "we_a_family/we_a_family/api"

func (router RouterGroup) ImageRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.POST("images", app.ImageUploadView)
}
