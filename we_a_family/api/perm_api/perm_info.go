package perm_api

import (
	"github.com/gin-gonic/gin"
	"we_a_family/we_a_family/service"
)

type PermissionApi struct {
}

func (PermissionApi) CreatePermissionView(ctx *gin.Context) {
	service.CreatePermissionService(ctx)
}

func (PermissionApi) DeletePermissionView(ctx *gin.Context) {
	service.DeletePermissionService(ctx)

}

func (PermissionApi) UpdatePermissionView(ctx *gin.Context) {
	service.UpdatePermissionService(ctx)

}

func (PermissionApi) FindAllPermissionView(ctx *gin.Context) {
	service.FindAllPermissionService(ctx)

}

func (PermissionApi) FindSelfPermissionView(ctx *gin.Context) {
	service.FindSelfPermissionService(ctx)

}
