package settings_api

import (
	"github.com/gin-gonic/gin"
	"we_a_family/we_a_family/utils"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	utils.FailwithCode(utils.SettingsError, c)
}
