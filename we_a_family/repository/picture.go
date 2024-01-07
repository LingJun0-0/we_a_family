package repository

import (
	"time"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/utils"
)

func FindOnePictureById(id int) (Models.Picture, error) {
	var p Models.Picture
	err := global.DB.Where(Models.Picture{Id: id}).Find(&p).Error
	return p, err
}

func FindAllPicture() ([]Models.Picture, error) {
	var pictures []Models.Picture
	err := global.DB.Model([]Models.Picture{}).Find(&pictures).Error
	return pictures, err
}

func InsertOnePicture(name, code, url string) (Models.Picture, error) {
	var p Models.Picture
	p = Models.Picture{
		Name:      name,
		Url:       url,
		Code:      code,
		CreatedAt: time.Now().Format(utils.Timestamp),
		UpdatedAt: time.Now().Format(utils.Timestamp)}
	err := global.DB.Create(&p).Error
	return p, err
}

func FindPicturesByNameAndCode(name, code string) (Models.Picture, error) {
	var picture Models.Picture
	err := global.DB.Model(&picture).
		Where(Models.Picture{Name: name, Code: code}).
		Find(&picture).Error
	if err != nil {
		return picture, err
	}

	return picture, nil
}

func FindTagsByPictureId(id int) {

}

func UpdateOnePictureById(id int, name, url string) error {
	err := global.DB.Model(&Models.Picture{}).
		Where(Models.Picture{Id: id}).
		Updates(Models.Picture{
			Name:      name,
			Url:       url,
			UpdatedAt: time.Now().Format(utils.Timestamp)}).Error
	return err
}

// UpdateOneTagPictureByPictureId    更新 tag_pictures tagId
func UpdateOneTagPictureByPictureId(pictureId, oldTagId, newTagId int) error {
	var picture Models.Picture
	err := global.DB.Model(&picture).Preload("Tags").Find(&picture, pictureId).Error
	if err != nil {
		return err
	}
	// 查找要删除的旧照片
	var oldTag Models.Tag
	err = global.DB.First(&oldTag, oldTagId).Error
	if err != nil {
		return err
	}
	// 查找要添加的新照片
	var newTag Models.Tag
	err = global.DB.First(&newTag, newTagId).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&picture).Association("Tags").Delete(&oldTag)
	if err != nil {
		return err
	}
	err = global.DB.Model(&picture).Association("Tags").Append(&newTag)
	if err != nil {
		return err
	}
	return nil

}

func DelOnePictureById(id int) error {
	err := global.DB.Delete(&Models.Picture{Id: id}).Error
	return err
}

// DelTagPictureByPictureId  删除 tag_pictures pictureId
func DelTagPictureByPictureId(pictureId int) error {
	var picture Models.Picture
	err := global.DB.Model(&picture).Preload("Tags").Find(&picture, pictureId).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&picture).Association("Tags").Delete(&picture.Tags)
	return err
}
