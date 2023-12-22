package repository

import (
	"fmt"
	"time"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/utils"
)

// FindOnePictureById 按照片id查信息
func FindOnePictureById(id int) (Models.Picture, error) {
	var p Models.Picture
	err := global.DB.Where("id = ?", id).Find(&p)
	if err.Error != nil {
		return p, err.Error
	}
	return p, err.Error
}

// FindAllPicture  查询所有照片数据
func FindAllPicture() []Models.Picture {
	var pictures []Models.Picture
	rows := global.DB.Find(&pictures)
	if rows.Error != nil {
		fmt.Printf("findsData failed err:%s\n", rows.Error)
		return nil
	}
	return pictures
}

// InsertOnePicture  跟据 name,MD5, url 插入一条数据,需要判断数据库里是否有同名或者同MD5照片
func InsertOnePicture(name, md5, url string) (err error) {
	var p Models.Picture
	err = global.DB.Where("name = ?", name).Or("code = ?", md5).First(&p).Error
	if err == nil {
		global.Log.Error("已存在相同照片")
		return

	} else {
		p = Models.Picture{Name: name, Url: url, Code: md5, CreatedAt: time.Now().Format(utils.Timestemp), UpdatedAt: time.Now().Format(utils.Timestemp)}
		res := global.DB.Create(&p)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

// UpdateOnePictureById 更新一行照片的路径
func UpdateOnePictureById(id int, name, url string) error {
	res := global.DB.Model(&Models.Picture{}).Where("id = ?", id).Select("name", "url").Updates(Models.Picture{Name: name, Url: url, UpdatedAt: time.Now().Format(utils.Timestemp)})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// DelOnePictureById 按id删除一行照片记录
func DelOnePictureById(id int) error {
	ret := global.DB.Delete(&Models.Picture{}, "id = ?", id)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
