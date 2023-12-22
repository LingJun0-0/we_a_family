package service

import (
	"github.com/gin-gonic/gin"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/repository"
	"we_a_family/we_a_family/utils"
)

func CreateTagService(ctx *gin.Context) {
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
	if member.Status >= int(utils.MemberStatusCode3) {
		var tag Models.Tag
		if err := ctx.ShouldBindJSON(&tag); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.InsertOneTagWithOwnerId(tag.Name, tag.Description, memberId.(int))
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		t, err := repository.FindTagWithOwnerIdAndName(memberId.(int), tag.Name)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(t, ctx)
		}
	} else {
		utils.FailwithMessage("用户无创建标签权限", ctx)
	}
}

func DeleteTagService(ctx *gin.Context) {
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
		var tag Models.Tag
		if err := ctx.ShouldBindJSON(&tag); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.DelOneTagById(tag.Id)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		t, err := repository.FindOneTagById(tag.Id)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(t, ctx)
		}
	} else {
		utils.FailwithMessage("用户无删除标签权限", ctx)
	}
}

func UpdateTagService(ctx *gin.Context) {
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
	if member.Status >= int(utils.MemberStatusCode3) {
		var tag Models.Tag
		if err := ctx.ShouldBindJSON(&tag); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		t, err := repository.FindOneTagById(tag.Id)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.UpdateOneTagById(t.Id, tag.Name, tag.Description, tag.Deleted)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		t, err = repository.FindTagWithOwnerIdAndName(memberId.(int), tag.Name)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(t, ctx)
		}
	} else {
		utils.FailwithMessage("用户无修改标签权限", ctx)
	}
}

func FindTagService(ctx *gin.Context) {
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
	if member.Status >= int(utils.MemberStatusCode1) {
		var tag Models.Tag
		if err := ctx.ShouldBindJSON(&tag); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		t, err := repository.FindTagByTagName(tag.Name)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(t, ctx)
		}
	} else {
		utils.FailwithMessage("用户无查找标签权限", ctx)
	}
}

func FindAllTagService(ctx *gin.Context) {
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
		t, err := repository.FindAllTag()
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(t, ctx)
		}
	} else {
		utils.FailwithMessage("用户无查找所有标签权限", ctx)
	}
}
