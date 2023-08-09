package routers

import (
	"we_a_family/we_a_family/api"
)

func (router RouterGroup) MemberRouter() {
	memberApi := api.ApiGroupApp.MemberApi
	//router.GET("login/", memberApi.MemberLoginInfoView)
	router.POST("login/", memberApi.MemberLoginInfoView)
	router.GET("memberList/", memberApi.MemberFindAll)
	router.POST("insert/", memberApi.InsertMemberView)
	router.PATCH("update/", memberApi.UpdateMemberView)
	router.DELETE("delete/", memberApi.DeleteMemberView)

	router.GET("/", memberApi.LoadMemberLoginHtml)

}
