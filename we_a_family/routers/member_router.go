package routers

import (
	"we_a_family/we_a_family/api"
)

func (router RouterGroup) MemberRouter() {
	memberApi := api.ApiGroupApp.MemberApi
	router.GET("memberlogin/", memberApi.MemberLoginInfoView)
	router.GET("findall/", memberApi.MemberFindAll)
}
