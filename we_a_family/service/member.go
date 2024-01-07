package service

import (
	"github.com/gin-gonic/gin"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/repository"
	"we_a_family/we_a_family/utils"
)

func MemberLoginService(ctx *gin.Context) {
	var user Models.Member
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	member, err := repository.LoginFindMember(user.Username, user.Password)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(member.Username, utils.Member)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	if perm.Code == 0 && perm.Description == utils.Member {
		utils.FailwithCode(utils.DeletedMember, ctx)
		return
	}
	if perm.Code >= int(utils.Reader) {
		token, err := GenToken(member.Username)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(token, ctx)
		}

	}

}

func RegisterMemberService(ctx *gin.Context) {
	var user Models.Member
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if _, err := repository.InsertOneMember(user.Username, user.Password); err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if err := repository.InsertOnePermission(user.Username, int(utils.Reader), utils.Member); err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if err := repository.InsertOnePermission(user.Username, int(utils.Writer), utils.Tag); err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if err := repository.InsertOnePermission(user.Username, int(utils.Writer), utils.Picture); err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if err := repository.InsertOnePermission(user.Username, int(utils.Reader), utils.Permission); err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	member, err := repository.LoginFindMember(user.Username, user.Password)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	} else {
		utils.OkwithData(member, ctx)
	}
}

// UpdateMemberSelfService 只允许修改密码
func UpdateMemberSelfService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	member, err := repository.FindOneMemberByUsername(memberId.(int))
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(member.Username, utils.Member)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if perm.Code >= int(utils.Reader) && perm.Code <= int(utils.Writer) {
		var user Models.Member
		if err := ctx.ShouldBindJSON(&user); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}

		if user.Password != "" {
			member.Password = user.Password
		} else {
			utils.FailwithMessage("自身不允许修改其他字段", ctx)
			return
		}

		if err := repository.UpdatePasswordByUsername(member.Username, member.Password); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}

		member, err = repository.FindOneMemberByUsername(memberId.(int))
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(member, ctx)
		}
	}

}

func MemberFindAllService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Member)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	if perm.Code == int(utils.Admin) {
		members, err := repository.FindAllMemberInfo()
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(members, ctx)
		}
	} else {
		utils.FailwithMessage("当前用户权限不足", ctx)
	}

}

func InsertMemberService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	member, err := repository.FindOneMemberByUsername(memberId.(int))
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(member.Username, utils.Member)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	if perm.Code == int(utils.Admin) {
		var user Models.Member
		if err := ctx.ShouldBindJSON(&user); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		if _, err := repository.InsertOneMember(user.Username, user.Password); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		if err := repository.InsertOnePermission(user.Username, int(utils.Admin), utils.Member); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		if err := repository.InsertOnePermission(user.Username, int(utils.Admin), utils.Tag); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		if err := repository.InsertOnePermission(user.Username, int(utils.Admin), utils.Picture); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		if err := repository.InsertOnePermission(user.Username, int(utils.Admin), utils.Permission); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		member, err := repository.LoginFindMember(user.Username, user.Password)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(member, ctx)
		}
	} else {
		utils.FailwithMessage("当前用户权限限制", ctx)
	}

}

func DeleteMemberService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Member)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	if perm.Code == int(utils.Admin) {
		var user Models.Member
		if err := ctx.ShouldBindJSON(&user); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		if user.Username == utils.GodUserName {
			utils.FailwithMessage("该id不允许通过本系统删除，请联系数据库管理员", ctx)
			return
		}
		_, err := repository.FindOneMemberByUsername(user.Username)
		if err != nil {
			utils.FailwithCode(utils.MemberDoesNotExist, ctx)
			return
		}

		err = repository.DelMemberPictureByUsername(user.Username)
		if err != nil {
			utils.FailwithMessage("删除关联照片出错"+err.Error(), ctx)
			return
		}

		err = repository.DelMemberTagByUsername(user.Username)
		if err != nil {
			utils.FailwithMessage("删除关联标签出错"+err.Error(), ctx)
			return
		}
		err = repository.DelPermissionsByResourceId(user.Username)
		if err != nil {
			utils.FailwithMessage("删除关联权限出错"+err.Error(), ctx)
			return
		}

		if err := repository.DelOneMemberByUsername(user.Username); err != nil {
			utils.FailwithCode(utils.DeleteError, ctx)
			return
		} else {

			utils.OkwithMessage("删除用户及其关联tags,pictures成功", ctx)
		}
	}

}
