package repository

import (
	"time"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/utils"
)

// FindOneTagById 按标签 id 查标签信息
func FindOneTagById(id int) (Models.Tag, error) {
	var t Models.Tag
	err := global.DB.Model(&t).Find(&t, id).Error
	return t, err
}

func FindPicturesByTagName(name string) ([]Models.Tag, error) {
	var tags []Models.Tag
	err := global.DB.Model(&tags).
		Preload("Pictures").
		Where("name like ?", "%"+name+"%").
		Find(&tags).Error
	return tags, err
}

func FindPicturesByTagId(id int) ([]Models.Picture, error) {
	var tag Models.Tag
	err := global.DB.Model(&tag).
		Preload("Pictures").
		Find(&tag, id).Error
	return tag.Pictures, err
}

func FindAllTagSelfByMemberId(username int) ([]Models.Tag, error) {
	var member Models.Member
	err := global.DB.Model(&member).
		Preload("Tags.Pictures").
		Find(&member, username).Error
	if err != nil {
		return nil, err
	}
	return member.Tags, nil
}

// FindAllTagInfo FindAllTag  查询所有标签数据
func FindAllTagInfo() ([]Models.Tag, error) {
	var tags []Models.Tag
	err := global.DB.Model(Models.Tag{}).
		Preload("Pictures.Tags").
		Find(&tags).Error
	return tags, err
}

func InsertOneTagWithNameAndDes(name, des string) (Models.Tag, error) {
	t := Models.Tag{
		Name:        name,
		Description: des,
		CreatedAt:   time.Now().Format(utils.Timestamp),
		UpdatedAt:   time.Now().Format(utils.Timestamp)}
	err := global.DB.Model(&t).Create(&t).Error
	return t, err
}

func UpdateOneTagById(id int, name, des string) error {
	err := global.DB.
		Where(Models.Tag{Id: id}).
		Updates(Models.Tag{
			Name:        name,
			Description: des,
			UpdatedAt:   time.Now().Format(utils.Timestamp)}).Error
	return err

}

func DelOneTagById(id int) error {
	err := global.DB.Delete(Models.Tag{Id: id}).Error
	return err
}

func DelTagByName(name string) error {
	var tag Models.Tag
	err := global.DB.Model(&tag).Where(Models.Tag{Name: name}).Delete(&tag).Error
	return err
}

// FindAllTagPictureByTagId   查询指定 tagId 所有 Picture
func FindAllTagPictureByTagId(tagId int) ([]Models.Picture, error) {
	var tag Models.Tag
	err := global.DB.Model(Models.Tag{}).
		Preload("Pictures.Tags").
		Where(Models.Tag{Id: tagId}).
		Find(&tag).Error
	if err != nil {
		return nil, err
	}
	var pictures []Models.Picture
	for _, picture := range tag.Pictures {
		pictures = append(pictures, picture)
	}
	return pictures, err
}

// FindOneTagPictureByTagIdAndPictureId  查询指定 tagId 指定 PictureId
func FindOneTagPictureByTagIdAndPictureId(tagId, pictureId int) (Models.Picture, error) {
	var tag Models.Tag
	var picture Models.Picture
	err := global.DB.Model(Models.Tag{}).
		Preload("Pictures.Tags").
		Where(Models.Tag{Id: tagId}).
		Find(&tag).Error
	if err != nil {
		return picture, err
	}

	for _, p := range tag.Pictures {
		if p.Id == pictureId {
			picture = p
		}
	}

	return picture, err
}

// InsertOneTagPictureWithTagIdAndPictureId  添加 tag_pictures tagId ,pictureId
func InsertOneTagPictureWithTagIdAndPictureId(tagId, pictureId int) error {
	var tag Models.Tag
	var p Models.Picture
	err := global.DB.Model(Models.Tag{}).
		Where(Models.Tag{Id: tagId}).
		Find(&tag).Error
	if err != nil {
		return err
	}
	err = global.DB.First(&p, pictureId).Error
	if err != nil {
		return err
	}
	// 添加关联关系
	err = global.DB.Model(&tag).Association("Pictures").Append(&p)
	if err != nil {
		return err
	}
	return err
}

// DelOneTagPictureByTagId  删除 tag_pictures username,pictureId
func DelOneTagPictureByTagId(tagId, pictureId int) error {
	var tag Models.Tag
	err := global.DB.Model(&tag).Preload("Pictures").Find(&tag, tagId).Error
	if err != nil {
		return err
	}
	// 查找要删除的照片
	var picture Models.Picture
	err = global.DB.First(&picture, pictureId).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&tag).Association("Pictures").Delete(&picture)
	if err != nil {
		return err
	}

	return err
}

// UpdateOneTagPictureByTagId   更新 tag_pictures pictureId
func UpdateOneTagPictureByTagId(tagId, oldPicturesId, newPicturesId int) error {
	var tag Models.Tag
	err := global.DB.Model(&tag).Preload("Pictures").Find(&tag, tagId).Error
	if err != nil {
		return err
	}
	// 查找要删除的旧照片
	var oldPicture Models.Picture
	err = global.DB.First(&oldPicture, oldPicturesId).Error
	if err != nil {
		return err
	}
	// 查找要添加的新照片
	var newPicture Models.Picture
	err = global.DB.First(&newPicture, newPicturesId).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&tag).Association("Pictures").Delete(&oldPicture)
	if err != nil {
		return err
	}
	err = global.DB.Model(&tag).Association("Pictures").Append(&newPicture)
	if err != nil {
		return err
	}
	return nil

}

// DelTagPictureByTagId  删除 tag_pictures tagId
func DelTagPictureByTagId(tagId int) error {
	var tag Models.Tag
	err := global.DB.Model(&tag).Preload("Pictures").Find(&tag, tagId).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&tag).Association("Pictures").Delete(&tag.Pictures)
	return err
}

// DelMemberTagByTagId  删除 tag_pictures tagId
func DelMemberTagByTagId(tagId int) error {
	var tag Models.Tag
	err := global.DB.Model(&tag).Preload("Members").Find(&tag, tagId).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&tag).Association("Members").Delete(&tag.Members)
	return err
}
