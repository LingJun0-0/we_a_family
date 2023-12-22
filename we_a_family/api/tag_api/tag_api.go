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

func (TagApi) DeleteTagView(ctx *gin.Context) {
	service.DeleteTagService(ctx)
}

func (TagApi) UpdateTagView(ctx *gin.Context) {
	service.UpdateTagService(ctx)
}

func (TagApi) FindTagView(ctx *gin.Context) {
	service.FindTagService(ctx)
}

func (TagApi) FindAllTagView(ctx *gin.Context) {
	service.FindAllTagService(ctx)
}
