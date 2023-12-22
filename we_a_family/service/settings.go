package service

import (
	"github.com/gin-gonic/gin"
	"we_a_family/we_a_family/utils"
)

func SettingsInfoService(ctx *gin.Context) {
	utils.OkwithMessage("网站关于选项的集合", ctx)
}
