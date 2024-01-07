package routers

import (
	"we_a_family/we_a_family/api"
	"we_a_family/we_a_family/service"
)

func (router RouterGroup) PermRouter() {
	permApi := api.GroupApp.PermissionApi
	router.Use(service.Auth)
	router.GET("permission", permApi.FindAllPermissionView)
	router.GET("permissionSelf", permApi.FindSelfPermissionView)
	router.POST("permission/creat", permApi.CreatePermissionView)
	router.PATCH("permission/update", permApi.UpdatePermissionView)
	router.DELETE("permission/delete", permApi.DeletePermissionView)

}
