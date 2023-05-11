package member_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/utils"
)

func (MemberApi) MemberLoginInfoView(ctx *gin.Context) {
	username := ctx.Param("username")
	username1, _ := strconv.Atoi(username)
	password := ctx.Param("password")
	member, err := Models.LoginFindMember(username1)
	if err != nil {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
	} else if member.Deleted {
		utils.FailwithCode(utils.DeletedMember, ctx)
	} else if member.Password == password {
		utils.OkwithData(member, ctx)
	} else {
		utils.FailwithCode(utils.NameOrPwdNotRight, ctx)
	}
}

func (MemberApi) MemberFindAll(ctx *gin.Context) {
	members := Models.FindsAllMember()
	for _, member := range members {
		if member.Deleted {
			utils.FailwithCode(utils.ListError, ctx)
		} else {
			utils.OkwithData(member, ctx)
		}
	}
}

func (MemberApi) InsertMemberView(ctx *gin.Context) {
	username := ctx.Param("username")
	username1, _ := strconv.Atoi(username)
	password := ctx.Param("password")
	member, err := Models.LoginFindMember(username1)
	if err != nil {
		err = Models.InsertOneMember(username1, password)
		if err != nil {
			utils.FailwithCode(utils.RegisterError, ctx)
		} else {
			member, _ := Models.LoginFindMember(username1)
			utils.OkwithData(member, ctx)
		}
	} else if member.Deleted {
		utils.FailwithCode(utils.DeletedMember, ctx)
	} else {
		utils.FailwithCode(utils.RegisterAgainError, ctx)
	}

}

func (MemberApi) UpdateMemberView(ctx *gin.Context) {
	username := ctx.Param("username")
	username1, _ := strconv.Atoi(username)
	password := ctx.Param("password")
	member, err := Models.LoginFindMember(username1)
	if err != nil {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
	} else if member.Password == password {
		id := member.Id
		changeUsername := ctx.Param("cusername")
		changeUsername1, _ := strconv.Atoi(changeUsername)
		changePassword := ctx.Param("cpwd")
		err := Models.UpdateOneMemberById(id, changeUsername1, changePassword, false)
		if err != nil {
			utils.FailwithCode(utils.ChangeError, ctx)
		} else {
			member, err = Models.FindOneMemberById(id)
			utils.OkwithData(member, ctx)
		}
	} else {
		utils.FailwithCode(utils.NameOrPwdNotRight, ctx)
	}
}

func (MemberApi) DeleteMemberView(ctx *gin.Context) {
	username := ctx.Param("username")
	username1, _ := strconv.Atoi(username)
	password := ctx.Param("password")
	member, err := Models.LoginFindMember(username1)
	if err != nil {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
	} else if member.Deleted {
		utils.FailwithCode(utils.DeletedMember, ctx)
	} else if member.Password == password {
		id := member.Id
		err := Models.DelOneMemberById(id)
		if err != nil {
			utils.FailwithCode(utils.DeleteError, ctx)
		} else {
			member, err = Models.FindOneMemberById(id)
			utils.OkwithData(member, ctx)
		}
	} else {
		utils.FailwithCode(utils.NameOrPwdNotRight, ctx)
	}
}
