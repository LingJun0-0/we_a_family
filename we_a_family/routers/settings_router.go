package routers

import (
	"we_a_family/we_a_family/api"
)

func (router RouterGroup) SettingsRouter() {
	settingApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", settingApi.SettingsInfoView)
}
