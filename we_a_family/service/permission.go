package service

import (
	"github.com/gin-gonic/gin"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/repository"
	"we_a_family/we_a_family/utils"
)

func CreatePermissionService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithMessage("当前账户非法", ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Permission)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if perm.Code == int(utils.Admin) {
		var perm Models.Permission
		if err := ctx.ShouldBindJSON(&perm); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		if perm.ResourceId < 0 || perm.Code < 0 || perm.Description == "" {
			utils.FailwithMessage("字段值不合法", ctx)
			return
		}
		isExist := repository.SamePermission(perm.ResourceId, perm.Code, perm.Description)
		if isExist {
			utils.FailwithMessage("该账户已有当前描述权限，请移步修改权限代数", ctx)
			return
		}
		err := repository.InsertOnePermission(perm.ResourceId, perm.Code, perm.Description)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		per, err := repository.GetOnePermissionByAllArguments(perm.ResourceId, perm.Code, perm.Description)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(per, ctx)
			return
		}

	}

}

func DeletePermissionService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithMessage("当前账户非法", ctx)
		return
	}
	perms, err := repository.GetPermissionsByResourceId(memberId.(int))
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	for _, permission := range perms {
		if permission.Code == int(utils.Admin) && permission.Description == utils.Permission {
			var perm Models.Permission
			if err := ctx.ShouldBindJSON(&perm); err != nil {
				utils.FailwithMessage(err.Error(), ctx)
				return
			}
			if perm.ResourceId < 0 || perm.Code < 0 || perm.Description == "" {
				utils.FailwithMessage("字段值不合法", ctx)
				return
			}
			err := repository.DeleteOnePermission(perm.ResourceId, perm.Code, perm.Description)
			if err != nil {
				utils.FailwithMessage(err.Error(), ctx)
				return
			} else {
				utils.OkwithData("删除成功", ctx)
			}
		}
	}
}

func UpdatePermissionService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithMessage("当前账户非法", ctx)
		return
	}

	var perms []Models.Permission
	perms, err := repository.GetPermissionsByResourceId(memberId.(int))
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	for _, permission := range perms {
		if permission.Code == int(utils.Admin) && permission.Description == utils.Permission {
			var perm Models.Permission
			if err := ctx.ShouldBindJSON(&perm); err != nil {
				utils.FailwithMessage(err.Error(), ctx)
				return
			}
			if perm.ResourceId < 0 || perm.Code < 0 || perm.Description == "" {
				utils.FailwithMessage("字段值不合法", ctx)
				return
			}
			if perm.Description == utils.Permission || perm.Description == utils.Tag ||
				perm.Description == utils.Member || perm.Description == utils.Picture {
				if perm.Code < int(utils.Black) || perm.Code > int(utils.Admin) {
					utils.FailwithMessage("权限代数不合法", ctx)
					return
				}
				isExist := repository.SamePermission(perm.ResourceId, perm.Code, perm.Description)
				if isExist {
					utils.FailwithMessage("存在相同权限，请移步创建", ctx)
					return
				}
				err := repository.UpdateOnePermission(perm.ResourceId, perm.Code, perm.Description)
				if err != nil {
					utils.FailwithMessage(err.Error(), ctx)
					return
				}
				p, err := repository.GetOnePermissionByAllArguments(perm.ResourceId, perm.Code, perm.Description)
				if err != nil {
					utils.FailwithMessage(err.Error(), ctx)
					return
				}
				if p.Code == perm.Code && p.ResourceId == perm.ResourceId && p.Description == perm.Description {
					utils.OkwithData(p, ctx)
				}
			} else {
				utils.FailwithMessage("权限描述不合法", ctx)
			}

		}
	}
}

func FindAllPermissionService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithMessage("当前账户非法", ctx)
		return
	}
	perms, err := repository.GetPermissionsByResourceId(memberId.(int))
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	for _, permission := range perms {
		if permission.Code == int(utils.Admin) && permission.Description == utils.Permission {
			perms, err := repository.GetAllPermissions()
			if err != nil {
				utils.FailwithMessage(err.Error(), ctx)
				return
			} else {
				utils.OkwithData(perms, ctx)
			}
		}
		if permission.ResourceId == utils.GodUserName && permission.Description == utils.Permission {

		}
	}
}

func FindSelfPermissionService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithMessage("当前账户非法", ctx)
		return
	}
	perms, err := repository.GetPermissionsByResourceId(memberId.(int))
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	utils.OkwithData(perms, ctx)
}
