package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/repository"
	"we_a_family/we_a_family/utils"
)

var (
	// WhiteImageList 上传图片后缀白名单
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

type FileResponse struct {
	FileName  string `json:"file_name"`  //文件名
	IsSuccess bool   `json:"is_success"` //是否成功
	Msg       string `json:"msg"`        //消息
}

// ImageUploadService 上传图片
func ImageUploadService(ctx *gin.Context) {
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

		form, err := ctx.MultipartForm()
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		// form表单里名字为images的文件列表
		filelist, ok := form.File["images"]
		if !ok {
			utils.FailwithMessage("不存在文件", ctx)
			return
		}

		// 检查 upload 是否已存在文件
		err = utils.DirectoryIfNotExists(global.Config.Upload.Path)
		if err != nil {
			utils.FailwithMessage("无法创建目录: "+err.Error(), ctx)
			return
		}

		var resList []FileResponse
		// 可一次上传多个图片
		for _, file := range filelist {
			// 检查文件合法性
			err := validateFile(file)
			if err != nil {
				resList = append(resList, FileResponse{
					FileName:  file.Filename,
					IsSuccess: false,
					Msg:       err.Error(),
				})
				continue
			}

			// 拼接成 Upload.Path/file.Filename
			filePath := path.Join(global.Config.Upload.Path, file.Filename)
			//判断图片大小
			size := float64(file.Size) / float64(1024*1024)
			if size >= float64(global.Config.Upload.Size) {
				resList = append(resList, FileResponse{
					FileName:  file.Filename,
					IsSuccess: false,
					Msg:       fmt.Sprintf("图片超过设定大小，当前大小为%.2fMB, 设定大小为：%d MB", size, global.Config.Upload.Size),
				})
				continue
			}
			// 计算字节数据的 MD5 值用来判别是否同一照片
			fileObj, err := file.Open()
			if err != nil {
				global.Log.Error(err.Error())
				return
			}
			byteData, err := io.ReadAll(fileObj)
			md5String := utils.Md5(byteData)
			// 保存照片名字，MD5，路径到数据库
			if err = repository.InsertOnePicture(file.Filename, md5String, global.Config.Upload.Path); err != nil {
				global.Log.Error(err.Error())
				return
			}

			//SaveUploadedFile()里有os.MkdirAll(),未存在文件夹会自动创建
			err = ctx.SaveUploadedFile(file, filePath)
			if err != nil {
				global.Log.Error(err)
				resList = append(resList, FileResponse{
					FileName:  file.Filename,
					IsSuccess: false,
					Msg:       err.Error(),
				})
				continue
			}
			resList = append(resList, FileResponse{
				FileName:  filePath,
				IsSuccess: true,
				Msg:       "上传成功",
			})

		}
		utils.OkwithData(resList, ctx)
	}

}

// ImageFindAllService  查看已上传文件夹下所有图片
func ImageFindAllService(ctx *gin.Context) {
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

		var resList []FileResponse
		files, err := os.ReadDir(global.Config.Upload.Path)
		if err != nil {
			global.Log.Error(err.Error())
			return
		}

		for _, file := range files {
			fileName := file.Name()
			filePath := path.Join(global.Config.Upload.Path, fileName)
			err = validateDownloadFile(file)
			if err != nil {
				resList = append(resList, FileResponse{
					FileName:  filePath,
					IsSuccess: false,
					Msg:       "非法数据",
				})
			} else {
				resList = append(resList, FileResponse{
					FileName:  filePath,
					IsSuccess: true,
					Msg:       "加载成功",
				})
			}

		}

		utils.OkwithData(resList, ctx)
	}
}

// ImageDownloadService  下载图片(根据url传回数据库id)
func ImageDownloadService(ctx *gin.Context) {
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

		// 检查 download 路径是否已存在,不存在就创建
		err := utils.DirectoryIfNotExists(global.Config.Download.Path)
		if err != nil {
			utils.FailwithMessage("无法创建目录: "+err.Error(), ctx)
			return
		}

		var resList []FileResponse
		// 获取下载照片id，如果没有获取上传目录中所有照片，有则下载单张照片
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			files, err := getFilesInUploadDirectory(ctx)
			if err != nil {
				global.Log.Error(err.Error())
				return
			}
			for _, file := range files {
				uploadPath := path.Join(global.Config.Upload.Path, file.Name())
				downloadPath := path.Join(global.Config.Download.Path, file.Name())
				err = copyFile(uploadPath, downloadPath)
				if err != nil {
					resList = append(resList, FileResponse{
						FileName:  file.Name(),
						IsSuccess: false,
						Msg:       err.Error(),
					})
				}

				resList = append(resList, FileResponse{
					FileName:  downloadPath,
					IsSuccess: true,
					Msg:       "下载成功",
				})

			}

		} else {
			p, err := repository.FindOnePictureById(id)
			files, err := ioutil.ReadDir(p.Url)
			if err != nil {
				global.Log.Error(err.Error())
				return
			}

			for _, file := range files {
				if !file.IsDir() {
					fileName := file.Name()
					filePath := path.Join(p.Url, fileName)
					if p.Name == fileName {
						// 计算字节数据的 MD5 值用来判别是否同一照片
						fileObj, err := os.Open(filePath)
						if err != nil {
							global.Log.Error(err.Error())
							return
						}
						byteData, err := io.ReadAll(fileObj)
						md5String := utils.Md5(byteData)
						if md5String == p.Code {
							DbPath := path.Join(p.Url, p.Name)
							downloadPath := path.Join(global.Config.Download.Path, file.Name())

							err = copyFile(DbPath, downloadPath)
							if err != nil {
								resList = append(resList, FileResponse{
									FileName:  file.Name(),
									IsSuccess: false,
									Msg:       err.Error(),
								})
							}

							resList = append(resList, FileResponse{
								FileName:  downloadPath,
								IsSuccess: true,
								Msg:       "下载成功",
							})
						}

					}

				}
			}
		}
		utils.OkwithData(resList, ctx)
	}

}

// ImageDeleteService   删除数据库图片数据(不删除upload文件)
func ImageDeleteService(ctx *gin.Context) {
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

		id, _ := strconv.Atoi(ctx.Param("id"))
		if err := repository.DelOnePictureById(id); err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithMessage("删除成功", ctx)
		}
	}
}

// ImageUpdateService 修改数据库数据
func ImageUpdateService(ctx *gin.Context) {
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
		var p Models.Picture
		err = ctx.ShouldBindJSON(&p)
		if err != nil {
			global.Log.Error(err.Error())
			return
		}
		picture, err := repository.FindOnePictureById(p.Id)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		//取得名称后缀
		oldExtension := path.Ext(picture.Name)
		newExtension := path.Ext(p.Name)

		if p.Name == "" {
			p.Name = picture.Name
		} else if newExtension == "" || !utils.InList(newExtension, WhiteImageList) {
			p.Name = p.Name + oldExtension
		}
		if p.Url == "" {
			p.Url = picture.Url
		}
		oldName := path.Join(picture.Url, picture.Name)
		newName := path.Join(p.Url, p.Name)
		if len(newName) > 50 {
			utils.FailwithMessage("文件名超长，请重新输入", ctx)
			return
		}
		// 对源文件重命名
		err = os.Rename(oldName, newName)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		err = repository.UpdateOnePictureById(picture.Id, p.Name, p.Url)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		}
		picture, err = repository.FindOnePictureById(p.Id)
		if err != nil {
			utils.FailwithMessage(err.Error(), ctx)
			return
		} else {
			utils.OkwithData(picture, ctx)
		}

	}

}

// validateFile 验证上传文件合法性
func validateFile(file *multipart.FileHeader) error {
	// 检查文件后缀合法性
	extension := strings.ToLower(path.Ext(file.Filename))
	if !utils.InList(extension, WhiteImageList) {
		return errors.New("非法文件")
	}

	// 检查文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		return fmt.Errorf("图片超过设定大小，当前大小为%.2fMB，设定大小为：%d MB", size, global.Config.Upload.Size)
	}

	return nil
}

// validateDownloadFile 验证下载文件合法性
func validateDownloadFile(file os.DirEntry) error {
	// 检查文件后缀合法性
	extension := strings.ToLower(path.Ext(file.Name()))
	if !file.IsDir() && !utils.InList(extension, WhiteImageList) {
		return errors.New("非法文件")
	}

	return nil
}

// getFilesInUploadDirectory 获取上传目录中的文件
func getFilesInUploadDirectory(ctx *gin.Context) ([]fs.DirEntry, error) {
	dir, err := os.Open(global.Config.Upload.Path)
	if err != nil {
		utils.FailwithMessage("无法打开目录: "+err.Error(), ctx)
		return nil, err
	}
	defer dir.Close()

	files, err := dir.ReadDir(0)
	if err != nil {
		utils.FailwithMessage("无法读取目录: "+err.Error(), ctx)
		return nil, err
	}

	return files, nil
}

// copyFile 复制文件
func copyFile(srcPath, destPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		global.Log.Error("保存图片失败")
		return err
	}

	return nil
}
