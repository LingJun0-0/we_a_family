package tag_api

import (
	"github.com/gin-gonic/gin"
	"we_a_family/we_a_family/service"
)

type TagApi struct {
}

func (TagApi) CreateTagView(ctx *gin.Context) {
	service.CreateTagService(ctx)
}

func (TagApi) DeleteSelfTagView(ctx *gin.Context) {
	service.DeleteSelfTagService(ctx)
}

func (TagApi) UpdateTagView(ctx *gin.Context) {
	service.UpdateTagService(ctx)
}

func (TagApi) FindTagView(ctx *gin.Context) {
	tagId := ctx.Query("tagId")
	tagName := ctx.Query("tagName")
	if tagName != "" {
		service.FindPictureServiceByTagName(ctx)
	}
	if tagId != "" {
		service.FindPicturesServiceByTagId(ctx)

	} else {
		service.FindTagServiceByMemberId(ctx)
	}
}

func (TagApi) FindAllTagView(ctx *gin.Context) {
	service.FindAllTagsService(ctx)
}

func (TagApi) DeleteTagView(ctx *gin.Context) {
	service.DeleteTag(ctx)
}
