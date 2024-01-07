package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/utils"
)

func LoginFindMember(username int, password string) (Models.Member, error) {
	var m Models.Member
	err := global.DB.Model(&m).Find(&m, username).Error
	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	if err != nil {
		return m, errors.New("密码错误")
	}
	return m, nil
}

// FindOneMemberByUsername FindOneMemberById 按用户id查用户信息
func FindOneMemberByUsername(username int) (Models.Member, error) {
	var m Models.Member
	err := global.DB.Model(&m).Find(&m, username).Error
	return m, err
}

// FindAllMemberInfo  查询所有用户数据
func FindAllMemberInfo() ([]Models.Member, error) {
	var members []Models.Member
	err := global.DB.Model(&members).
		Preload("Tags").Preload("Pictures").
		Find(&members).Error
	return members, err
}

// InsertOneMember 跟据输入的用户名和密码插入一条新数据
func InsertOneMember(username int, password string) (Models.Member, error) {
	var m Models.Member
	// 生成密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return m, err
	}
	m = Models.Member{
		Username:  username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().Format(utils.Timestamp),
		UpdatedAt: time.Now().Format(utils.Timestamp)}
	err = global.DB.Model(&m).Create(&m).Error
	return m, err
}

// UpdatePasswordByUsername 更新密码
func UpdatePasswordByUsername(username int, password string) error {
	// 生成密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = global.DB.
		Where(Models.Member{Username: username}).
		Updates(Models.Member{
			Password:  string(hashedPassword),
			UpdatedAt: time.Now().Format(utils.Timestamp)}).Error
	return err
}

// DelOneMemberByUsername  按id删除一个用户
func DelOneMemberByUsername(username int) error {
	var m Models.Member
	err := global.DB.Model(&m).Delete(&m, username).Error
	return err
}

// FindAllTagsByUsername  查询当前用户所有tag
func FindAllTagsByUsername(username int) ([]Models.Tag, error) {
	var member Models.Member
	err := global.DB.Model(&member).
		Preload("Tags").
		Find(&member, username).Error

	return member.Tags, err
}

// FindOneMemberTagByUsernameAndTagId  查询当前用户指定tagId
func FindOneMemberTagByUsernameAndTagId(username, tagId int) (Models.Tag, error) {
	var member Models.Member
	var tag Models.Tag
	err := global.DB.Model(&member).Preload("Tags").
		Find(&member, username).Error
	if err != nil {
		return tag, err
	}

	for _, t := range member.Tags {
		if t.Id == tagId {
			tag = t
		}
	}

	return tag, nil
}

// InsertOneMemberTagByUsernameAndTagId  添加member_tags username,tagId
func InsertOneMemberTagByUsernameAndTagId(username, tagId int) error {
	var m Models.Member
	var t Models.Tag
	err := global.DB.Model(&m).Preload("Tags").Find(&m, username).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&t).Find(&t, tagId).Error
	if err != nil {
		return err
	}
	// 添加关联关系
	err = global.DB.Model(&m).Association("Tags").Append(&t)
	if err != nil {
		return err
	}
	return nil
}

// UpdateMemberTagByUsernameWithOldTagIdAndNewTagId 更新member_tags tagId
func UpdateMemberTagByUsernameWithOldTagIdAndNewTagId(username, oldTagId, newTagId int) error {
	var m Models.Member
	err := global.DB.Model(&m).Preload("Tags").Find(&m, username).Error
	if err != nil {
		return err
	}
	// 查找要删除的旧标签
	var oldTag Models.Tag
	err = global.DB.First(&oldTag, oldTagId).Error
	if err != nil {
		return err
	}
	// 查找要添加的新标签
	var newTag Models.Tag
	err = global.DB.First(&newTag, newTagId).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&m).Association("Tags").Delete(&oldTag)
	if err != nil {
		return err
	}
	err = global.DB.Model(&m).Association("Tags").Append(&newTag)
	if err != nil {
		return err
	}
	return nil
}

// DelOneMemberTagByUsernameAndTagId 删除member_tags username,tagId
func DelOneMemberTagByUsernameAndTagId(username, tagId int) error {
	var m Models.Member
	err := global.DB.Model(&m).Preload("Tags").Find(&m, username).Error
	if err != nil {
		return err
	}
	// 查找要删除的旧标签
	var tag Models.Tag
	err = global.DB.Model(&tag).First(&tag, tagId).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&m).Association("Tags").Delete(&tag)
	if err != nil {
		return err
	}
	return nil
}

func DelMemberTagByUsername(username int) error {
	var member Models.Member
	err := global.DB.Model(&member).Preload("Tags").Find(&member, username).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&member).Association("Tags").Delete(&member.Tags)

	return err
}

// FindAllMemberPictureByUsername  查询指定 username 所有 Picture
func FindAllMemberPictureByUsername(username int) ([]Models.Picture, error) {
	var member Models.Member
	err := global.DB.Model(Models.Member{}).
		Preload("Pictures.Tags").
		Where(Models.Member{Username: username}).
		Find(&member).Error

	return member.Pictures, err
}

// FindOneMemberPictureByUsernameAndPictureId  查询指定 username 指定 PictureId
func FindOneMemberPictureByUsernameAndPictureId(username, pictureId int) (Models.Picture, error) {
	var member Models.Member
	var picture Models.Picture
	err := global.DB.Model(Models.Member{}).
		Preload("Pictures.Tags").
		Where(Models.Member{Username: username}).
		Find(&member).Error
	if err != nil {
		return picture, err
	}

	for _, p := range member.Pictures {
		if p.Id == pictureId {
			picture = p
		}
	}

	return picture, err
}

// InsertOneMemberPictureByUsernameAndPictureId InsertOneMemberPictureByUsername  添加 tag_pictures username ,pictureId
func InsertOneMemberPictureByUsernameAndPictureId(username, pictureId int) error {
	var member Models.Member
	var p Models.Picture
	err := global.DB.Model(&member).Find(&member, username).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&p).First(&p, pictureId).Error
	if err != nil {
		return err
	}
	// 添加关联关系
	err = global.DB.Model(&member).Association("Pictures").Append(&p)
	if err != nil {
		return err
	}
	return err
}

// UpdateMemberPictureByUsernameWithOldPictureIdAndNewPictureId 更新 tag_pictures pictureId
func UpdateMemberPictureByUsernameWithOldPictureIdAndNewPictureId(username, oldPicturesId, newPicturesId int) error {
	var member Models.Member
	err := global.DB.Model(&member).Preload("Pictures").Find(&member, username).Error
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
	err = global.DB.Model(&member).Association("Pictures").Delete(&oldPicture)
	if err != nil {
		return err
	}
	err = global.DB.Model(&member).Association("Pictures").Append(&newPicture)
	if err != nil {
		return err
	}
	return nil

}

// DelOneMemberPictureByUsernameAndPictureId 删除 tag_pictures username,pictureId
func DelOneMemberPictureByUsernameAndPictureId(username, pictureId int) error {
	var member Models.Member
	err := global.DB.Model(&member).Preload("Pictures").Find(&member, username).Error
	if err != nil {
		return err
	}
	// 查找要删除的照片
	var picture Models.Picture
	err = global.DB.First(&picture, pictureId).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&member).Association("Pictures").Delete(&picture)
	if err != nil {
		return err
	}

	return err
}

func DelMemberPictureByUsername(username int) error {
	var member Models.Member
	err := global.DB.Model(&member).Preload("Pictures").Find(&member, username).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&member).Association("Pictures").Delete(&member.Pictures)

	return err
}
