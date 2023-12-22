package routers

import (
	"we_a_family/we_a_family/api"
	"we_a_family/we_a_family/service"
)

func (router RouterGroup) SettingsRouter() {
	settingApi := api.GroupApp.SettingsApi
	router.Use(service.Auth)
	router.GET("settings", settingApi.SettingsInfoView)
}
