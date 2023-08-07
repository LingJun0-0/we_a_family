package image_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"path"
	"strings"
	"we_a_family/we_a_family/global"
	"we_a_family/we_a_family/utils"
)

var (
	//上传图片后缀白名单
	WhiteImageList = []string{
		".jpg",
		".png",
		".jpeg",
		".ico",
		".tiff",
		".gif",
		".svg",
		".webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  //文件名
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //消息
}

// ImageUploadView 上传单个图片
func (ImageApi) ImageUploadView(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	filelist, ok := form.File["images"]
	if !ok {
		utils.FailwithMessage("不存在文件", ctx)
		return
	}

	var resList []FileUploadResponse

	for _, file := range filelist {
		nameList := path.Ext(file.Filename)
		suffix := strings.ToLower(nameList)
		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}

		filePath := path.Join(global.Config.Upload.Path, file.Filename)
		//判断图片大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片超过设定大小，当前大小为%.2fMB, 设定大小为：%d MB", size, global.Config.Upload.Size),
			})
			continue
		}

		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		md5String := utils.Md5(byteData)
		fmt.Println(md5String)

		//SaveUploadedFile()里有os.MkdirAll()的,未存在文件夹会自动创建
		err = ctx.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})

	}
	utils.OkwithData(resList, ctx)

}
