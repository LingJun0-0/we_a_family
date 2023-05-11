package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"we_a_family/we_a_family/global"
	"we_a_family/we_a_family/utils"
)

type Member struct {
	Id        int    `gorm:"primary_key"`
	Username  int    `gorm:"not null;unique_index"`
	Password  string `gorm:"type:varchar(255);not null"`
	Status    int
	Deleted   bool   `gorm:"column:deleted;type:tinyint(1)"`
	CreatedAt string `gorm:"column:created_at;type:datetime"`
	UpdatedAt string `gorm:"column:updated_at;type:datetime"`
	Tag       []Tag  `gorm:"foreignKey:OwnerId"`
	Perm      []Perm `gorm:"foreignKey:UserId"`
}

//第一次查用户按用户名
func LoginFindMember(username int) (Member, error) {
	var m Member
	err := global.DB.Where("username = ?", username).First(&m)
	if err.Error != nil {
		return m, err.Error
	}
	return m, err.Error
}

//按用户id查用户信息
func FindOneMemberById(id int) (Member, error) {
	var m Member
	err := global.DB.Where("id = ?", id).First(&m)
	if err.Error != nil {
		return m, err.Error
	}
	return m, err.Error
}

// 查询所有用户数据
func FindsAllMember() []Member {
	var members []Member
	rows := global.DB.Find(&members)
	if rows.Error != nil {
		fmt.Printf("findsData failed err:%s\n", rows.Error)
		return nil
	}
	return members
}

// 跟据输入的用户名和密码插入一条新数据
func InsertOneMember(username int, password string) (err error) {
	m := Member{Username: username, Password: password, Status: 1, Deleted: false, CreatedAt: time.Now().Format(utils.Timestemp), UpdatedAt: time.Now().Format(utils.Timestemp)}
	res := global.DB.Create(&m)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 更新一个用户的用户名，密码，delete状态
func UpdateOneMemberById(id, username int, password string, deleted bool) error {
	res := global.DB.Model(&Member{}).Where("id = ?", id).Select("username", "password", "deleted", "updated_at").Updates(Member{Username: username, Password: password, Deleted: deleted, UpdatedAt: time.Now().Format(utils.Timestemp)})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 按id删除一个用户
func DelOneMemberById(id int) error {
	ret := global.DB.Model(&Member{}).Where("id = ?", id).Select("deleted").Updates(Member{Deleted: true})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
