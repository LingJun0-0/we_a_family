package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"we_a_family/we_a_family/global"
)

type Member struct {
	Id         int    `gorm:"primary_key"`
	Username   string `gorm:"not null"`
	Password   string `gorm:"not null"`
	Status     int    `gorm:"not null"`
	Created_at string `gorm:"not null"`
	Updated_at string `gorm:"not null"`
	Deleted    bool   `gorm:"false"`
}

//第一次查用户按用户名
func FindOneMember(username string) (Member, error) {
	var m Member
	err := global.DB.Where("username = ?", username).First(&m)
	if err.Error != nil {
		fmt.Printf("findone member failed err :%s\n", err.Error)
	}
	fmt.Printf("findone member info %v\n", m)
	return m, err.Error
}

//按用户id查用户信息
func FindOneMemberUseId(id int) (Member, error) {
	var m Member
	err := global.DB.Where("id = ?", id).First(&m)
	if err.Error != nil {
		fmt.Printf("findone member failed err :%s\n", err.Error)
	}
	fmt.Printf("findone member info %v\n", m)
	return m, err.Error
}

// 查询所有用户数据
func FindsAllMember() []Member {
	var s []Member
	rows := global.DB.Find(&s)
	if rows.Error != nil {
		fmt.Printf("findsData failed err:%s\n", rows.Error)
		return nil
	}
	return s
}

// 跟据输入的用户名和密码插入一条新数据
func InsertOneMember(username string, password string) (err error) {
	m := Member{Username: username, Password: password, Status: 1, Deleted: false, Created_at: time.Now().Format(global.TimeFormat), Updated_at: time.Now().Format(global.TimeFormat)}
	res := global.DB.Create(&m)
	if res.Error != nil {
		fmt.Printf("Create insert failed err:%s\n", res.Error)
		return res.Error
	}
	ret := global.DB.Where("username = ?", username).First(&m)
	if ret.Error != nil {
		fmt.Printf("exec insert failed err:%s\n", ret.Error)
		return ret.Error
	}
	fmt.Printf("insert data id is : %d\n", m.Id)
	return nil
}

// 更新一个用户的用户名，密码，delete状态
func UpdateOneMember(id int, username, password string, delete bool) error {
	res := global.DB.Model(&Member{}).Where("id = ?", id).Select("username", "password", "delete", "updated_at").Updates(Member{Username: username, Password: password, Deleted: delete, Updated_at: time.Now().Format(global.TimeFormat)})
	if res.Error != nil {
		fmt.Printf("save failed err:%s\n", res.Error)
		return res.Error
	}
	fmt.Printf("update success rows:%d\n", res.RowsAffected)
	return nil
}

// 按id删除一个用户
func DelOneMember(id int) error {
	ret := global.DB.Model(&Member{}).Where("id = ?", id).Select("deleted").Updates(Member{Deleted: true})
	if ret.Error != nil {
		fmt.Printf("save failed err:%s\n", ret.Error)
		return ret.Error
	}
	fmt.Printf("update success rows:%d\n", ret.RowsAffected)
	return nil
}