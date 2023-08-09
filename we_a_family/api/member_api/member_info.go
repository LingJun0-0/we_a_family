package member_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/utils"
)

func (MemberApi) LoadMemberLoginHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user/login.html", gin.H{
		"zhuce": "注册成功请登录",
	})

}

func (MemberApi) MemberLoginInfoView(ctx *gin.Context) {
	username, _ := strconv.Atoi(ctx.PostForm("username"))
	password := ctx.PostForm("password")
	member, err := Models.LoginFindMember(username)
	if err != nil {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
	} else if member.Deleted {
		utils.FailwithCode(utils.DeletedMember, ctx)
	} else if member.Password == password {
		//utils.OkwithData(member, ctx)
		ctx.HTML(http.StatusOK, "user/index.html", gin.H{
			"m_id":         member.Id,
			"m_username":   member.Username,
			"m_pwd":        member.Password,
			"m_created_at": member.CreatedAt,
			"m_updated_at": member.UpdatedAt,
			"m_deleted_at": member.Deleted,
		})
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
	username, _ := strconv.Atoi(ctx.Param("username"))
	password := ctx.Param("password")
	member, err := Models.LoginFindMember(username)
	if err != nil {
		err = Models.InsertOneMember(username, password)
		if err != nil {
			utils.FailwithCode(utils.RegisterError, ctx)
		} else {
			member, _ := Models.LoginFindMember(username)
			utils.OkwithData(member, ctx)
		}
	} else if member.Deleted {
		utils.FailwithCode(utils.DeletedMember, ctx)
	} else {
		utils.FailwithCode(utils.RegisterAgainError, ctx)
	}

}

func (MemberApi) UpdateMemberView(ctx *gin.Context) {
	username, _ := strconv.Atoi(ctx.Param("username"))
	password := ctx.Param("password")
	member, err := Models.LoginFindMember(username)
	if err != nil {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
	} else if member.Password == password {
		id := member.Id
		changeUsername, _ := strconv.Atoi(ctx.Param("cusername"))
		changePassword := ctx.Param("cpwd")
		err := Models.UpdateOneMemberById(id, changeUsername, changePassword, false)
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
	username, _ := strconv.Atoi(ctx.Param("username"))
	password := ctx.Param("password")
	member, err := Models.LoginFindMember(username)
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
