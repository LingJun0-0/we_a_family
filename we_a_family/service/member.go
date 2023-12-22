package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
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
	member, err := repository.LoginFindMember(user.Username)

	if err != nil {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	} else if member.Password != user.Password {
		utils.FailwithCode(utils.PwdNotRight, ctx)
		return
	} else if member.Deleted {
		utils.FailwithCode(utils.DeletedMember, ctx)
		return
	} else {
		token, err := GenToken(member.Id)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
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
	if _, err := repository.LoginFindMember(user.Username); err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if err := repository.InsertOneMember(user.Username, user.Password); err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	member, _ := repository.LoginFindMember(user.Username)
	utils.OkwithData(member, ctx)
}

func UpdateMemberSelfService(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var user Models.Member
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	member, err := repository.FindOneMemberById(id)
	if err != nil {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	if err := repository.UpdateOneMemberById(id, user.Status, user.Username, user.Password, user.Deleted); err != nil {
		utils.FailwithCode(utils.ChangeError, ctx)
		return
	}
	member, err = repository.FindOneMemberById(id)
	utils.OkwithData(member, ctx)
}

func MemberFindAllService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	member, err := repository.FindOneMemberById(memberId.(int))
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if member.Status == int(utils.MemberStatusCode5) {
		members := repository.FindAllMember()
		utils.OkwithData(members, ctx)
	}
}

func InsertMemberService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	member, err := repository.FindOneMemberById(memberId.(int))
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if member.Status >= int(utils.MemberStatusCode4) {
		var user Models.Member
		if err := ctx.ShouldBindJSON(&user); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		if _, err := repository.LoginFindMember(user.Username); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		if err := repository.InsertOneMember(user.Username, user.Password); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		member, _ := repository.LoginFindMember(user.Username)
		utils.OkwithData(member, ctx)
	}
}

func UpdateMemberService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	member, err := repository.FindOneMemberById(memberId.(int))
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if member.Status >= int(utils.MemberStatusCode4) {
		var user Models.Member
		if err := ctx.ShouldBindJSON(&user); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		member, err := repository.FindOneMemberById(user.Id)
		if err != nil {
			utils.FailwithCode(utils.MemberDoesNotExist, ctx)
			return
		}
		if err := repository.UpdateOneMemberById(user.Id, user.Status, user.Username, user.Password, user.Deleted); err != nil {
			utils.FailwithCode(utils.ChangeError, ctx)
			return
		}
		member, err = repository.FindOneMemberById(user.Id)
		utils.OkwithData(member, ctx)
	}
}

func DeleteMemberService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	member, err := repository.FindOneMemberById(memberId.(int))
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if member.Status >= int(utils.MemberStatusCode4) {
		var user Models.Member
		if err := ctx.ShouldBindJSON(&user); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		member, err := repository.FindOneMemberById(user.Id)
		if err != nil {
			utils.FailwithCode(utils.MemberDoesNotExist, ctx)
			return
		}
		if err := repository.DelOneMemberById(user.Id); err != nil {
			utils.FailwithCode(utils.DeleteError, ctx)
		} else {
			member, err = repository.FindOneMemberById(user.Id)
			utils.OkwithData(member, ctx)
		}
	}
}
