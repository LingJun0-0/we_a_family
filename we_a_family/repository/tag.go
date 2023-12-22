package repository

import (
	"time"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/utils"
)

// FindOneTagById 按标签 id 查标签信息
func FindOneTagById(Id int) (Models.Tag, error) {
	var t Models.Tag
	err := global.DB.Where("id = ?", Id).Find(&t)
	if err.Error != nil {
		return t, err.Error
	}
	return t, err.Error
}

// FindTagWithOwnerIdAndName   按标签 name,ownerId 查标签信息
func FindTagWithOwnerIdAndName(ownerId int, name string) (Models.Tag, error) {
	var t Models.Tag
	err := global.DB.Where("name = ? AND owner_id = ?", name, ownerId).Find(&t)
	if err.Error != nil {
		return t, err.Error
	}
	return t, err.Error
}

func FindTagByTagName(name string) ([]Models.Tag, error) {
	var t []Models.Tag
	err := global.DB.Where("name = ?", name).Find(&t)
	if err.Error != nil {
		return t, err.Error
	}
	return t, err.Error
}

// FindAllTag  查询所有标签数据
func FindAllTag() ([]Models.Tag, error) {
	var tags []Models.Tag
	rows := global.DB.Find(&tags)
	if rows.Error != nil {
		global.Log.Errorf("findsData failed err:%s\n", rows.Error)
		return nil, rows.Error
	}
	return tags, nil
}

// InsertOneTagWithOwnerId   跟据 name,des, ownerId 插入一条数据,需要判断数据库里是否有同名和同owner标签
func InsertOneTagWithOwnerId(name, des string, ownerId int) (err error) {
	var t Models.Tag
	err = global.DB.Where(&Models.Tag{
		Name:    name,
		OwnerID: ownerId}).
		First(&t).Error
	if err == nil {
		global.Log.Error("创建失败，该用户已存在相同标签")
		return
	} else {
		t = Models.Tag{
			Name:        name,
			Description: des,
			OwnerID:     ownerId,
			Deleted:     false,
			CreatedAt:   time.Now().Format(utils.Timestemp),
			UpdatedAt:   time.Now().Format(utils.Timestemp)}
		res := global.DB.Create(&t)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

// UpdateOneTagByOwnerId 更新标签的 name, des(描述)
func UpdateOneTagById(Id int, name, des string, delete bool) error {
	res := global.DB.Where("id = ? ", Id).
		Select("name", "description", "deleted").
		Updates(Models.Tag{Name: name, Description: des, Deleted: delete, UpdatedAt: time.Now().Format(utils.Timestemp)})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// DelOneTagById 按id删除一行标签记录
func DelOneTagById(id int) error {
	ret := global.DB.Model(Models.Tag{}).
		Where("id = ?", id).
		Updates(Models.Tag{Deleted: true})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
