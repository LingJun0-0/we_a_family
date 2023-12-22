package member_api

import (
	"github.com/gin-gonic/gin"
	"we_a_family/we_a_family/service"
)

type MemberApi struct {
}

func (MemberApi) MemberLoginView(ctx *gin.Context) {
	service.MemberLoginService(ctx)
}

func (MemberApi) MemberFindAllView(ctx *gin.Context) {
	service.MemberFindAllService(ctx)
}

func (MemberApi) RegisterMemberView(ctx *gin.Context) {
	service.RegisterMemberService(ctx)
}

func (MemberApi) InsertMemberView(ctx *gin.Context) {
	service.InsertMemberService(ctx)
}

func (MemberApi) UpdateMemberSelfView(ctx *gin.Context) {
	service.UpdateMemberSelfService(ctx)
}

func (MemberApi) DeleteMemberView(ctx *gin.Context) {
	service.DeleteMemberService(ctx)
}

func (MemberApi) UpdateMemberView(ctx *gin.Context) {
	service.UpdateMemberService(ctx)

}
