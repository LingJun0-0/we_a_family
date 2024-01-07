package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/repository"
	"we_a_family/we_a_family/utils"
)

func CreateTagService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithMessage("当前账户非法", ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Tag)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	if perm.Code >= int(utils.Writer) && perm.Code <= int(utils.Admin) {
		var tag Models.Tag
		if err := ctx.ShouldBindJSON(&tag); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		t, err := repository.InsertOneTagWithNameAndDes(tag.Name, tag.Description)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.InsertOneMemberTagByUsernameAndTagId(memberId.(int), t.Id)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		t, err = repository.FindOneTagById(t.Id)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(t, ctx)
		}
	}

}

func DeleteSelfTagService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Tag)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	if perm.Code >= int(utils.Writer) && perm.Code <= int(utils.Admin) {
		tagId, err := strconv.Atoi(ctx.Query("tagId"))
		if err != nil {
			utils.FailwithMessage("传入标签参数不合法", ctx)
			return
		}
		t, err := repository.FindOneTagById(tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		_, err = repository.FindOneMemberTagByUsernameAndTagId(memberId.(int), tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.DelMemberTagByTagId(tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.DelTagPictureByTagId(tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.DelOneTagById(t.Id)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}

	}

}

func DeleteTag(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Tag)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	if perm.Code == int(utils.Admin) {
		tagId, err := strconv.Atoi(ctx.Query("tagId"))
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		_, err = repository.FindOneTagById(tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.DelOneMemberTagByUsernameAndTagId(memberId.(int), tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.DelTagPictureByTagId(tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.DelOneTagById(tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}

	} else {
		utils.FailwithMessage("当前用户权限不足", ctx)
	}

}

func UpdateTagService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Tag)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	if perm.Code >= int(utils.Writer) && perm.Code <= int(utils.Admin) {

		tagId, err := strconv.Atoi(ctx.Query("tagId"))
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}

		t, err := repository.FindOneTagById(tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		var tag Models.Tag
		if err := ctx.ShouldBindJSON(&tag); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		if tag.Name != "" {
			t.Name = tag.Name
		}
		if tag.Description != "" {
			t.Description = tag.Description
		}

		err = repository.UpdateOneTagById(tagId, t.Name, t.Description)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}

		t, err = repository.FindOneTagById(tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(t, ctx)
		}

	}
}

func FindPictureServiceByTagName(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Tag)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}

	if perm.Code >= int(utils.Reader) {
		tagName := ctx.Query("tagName")
		tags, err := repository.FindPicturesByTagName(tagName)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			//ctx.HTML(http.StatusOK, "", tags)
			utils.OkwithData(tags, ctx)
		}
	}

}

func FindAllTagsService(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Tag)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if perm.Code == int(utils.Admin) {
		t, err := repository.FindAllTagInfo()
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		utils.OkwithData(t, ctx)
	}

}

func FindTagServiceByMemberId(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Tag)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if perm.Code >= int(utils.Owner) && perm.Code <= int(utils.Admin) {
		tags, err := repository.FindAllTagSelfByMemberId(memberId.(int))
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		utils.OkwithData(tags, ctx)

	}

}

func FindPicturesServiceByTagId(ctx *gin.Context) {
	memberId, ok := ctx.Get("memberId")
	if !ok {
		utils.FailwithCode(utils.MemberDoesNotExist, ctx)
		return
	}
	perm, err := repository.GetOnePermissionByResourceIdAndDescription(memberId.(int), utils.Tag)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	if perm.Code >= int(utils.Owner) {
		tagId, err := strconv.Atoi(ctx.Query("tagId"))
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		pictures, err := repository.FindPicturesByTagId(tagId)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		utils.OkwithData(pictures, ctx)

	}

}
