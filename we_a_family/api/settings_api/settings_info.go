package settings_api

import (
	"github.com/gin-gonic/gin"
	"we_a_family/we_a_family/service"
)

type SettingsApi struct {
}

func (SettingsApi) SettingsInfoView(ctx *gin.Context) {
	service.SettingsInfoService(ctx)
}
