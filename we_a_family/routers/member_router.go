package routers

import (
	"we_a_family/we_a_family/api"
)

func (router RouterGroup) MemberRouter() {
	memberApi := api.ApiGroupApp.MemberApi
	router.GET("login/", memberApi.MemberLoginInfoView)
	router.GET("memberlist/", memberApi.MemberFindAll)
	router.GET("insert/:username/:password", memberApi.InsertMemberView)
}
