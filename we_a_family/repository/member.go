package repository

import (
	"errors"
	"fmt"
	"time"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/utils"
)

// LoginFindMember 第一次查用户按用户名
func LoginFindMember(username string) (Models.Member, error) {
	var m Models.Member
	err := global.DB.Where("username = ?", username).Find(&m)
	if err.Error != nil {
		return m, err.Error
	}
	return m, err.Error
}

// FindOneMemberById 按用户id查用户信息
func FindOneMemberById(id int) (Models.Member, error) {
	var m Models.Member
	err := global.DB.Where("id = ?", id).Find(&m)
	if err.Error != nil {
		return m, err.Error
	}
	return m, err.Error
}

// FindAllMember 查询所有用户数据
func FindAllMember() []Models.Member {
	var members []Models.Member
	//rows := global.DB.Find(&members)
	rows := global.DB.Where("deleted = 0").Find(&members)
	if rows.Error != nil {
		fmt.Printf("findsData failed err:%s\n", rows.Error)
		return nil
	}
	return members
}

// InsertOneMember 跟据输入的用户名和密码插入一条新数据
func InsertOneMember(username string, password string) (err error) {
	var m Models.Member
	err = global.DB.
		Where("username = ?", username).
		First(&m).Error
	if err == nil {
		return errors.New("已存在同账户名")
	} else {
		m = Models.Member{Username: username, Password: password, Status: 1, Deleted: false, CreatedAt: time.Now().Format(utils.Timestemp), UpdatedAt: time.Now().Format(utils.Timestemp)}
		res := global.DB.Create(&m)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

// UpdateOneMemberById 更新一个用户的用户名，密码，delete,状态
func UpdateOneMemberById(id, status int, username string, password string, deleted bool) error {
	res := global.DB.
		Where("id = ?", id).
		Select("username", "password", "status", "deleted", "updated_at").
		Updates(Models.Member{Username: username, Password: password, Status: status, Deleted: deleted, UpdatedAt: time.Now().Format(utils.Timestemp)})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// DelOneMemberById 按id删除一个用户
func DelOneMemberById(id int) error {
	ret := global.DB.
		Where("id = ?", id).
		Updates(Models.Member{Deleted: true})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
