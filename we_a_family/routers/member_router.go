package routers

import (
	"we_a_family/we_a_family/api"
	"we_a_family/we_a_family/service"
)

func (router RouterGroup) MemberRouter() {
	memberApi := api.GroupApp.MemberApi
	router.GET("user/login", memberApi.MemberLoginView)
	router.POST("user/register", memberApi.RegisterMemberView)

	//路由鉴权
	router.Use(service.Auth)
	router.PATCH("user/updateSelf", memberApi.UpdateMemberSelfView)
	router.GET("user/memberList", memberApi.MemberFindAllView)
	router.POST("user/insert", memberApi.InsertMemberView)
	router.DELETE("user/delete", memberApi.DeleteMemberView)

}
