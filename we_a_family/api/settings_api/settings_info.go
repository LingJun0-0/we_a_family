package settings_api

import (
	"github.com/gin-gonic/gin"
	"we_a_family/we_a_family/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailwithCode(res.SettingsError, c)
}
