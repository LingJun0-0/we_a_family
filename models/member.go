package Models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"we_a_family/global"
)

type Member struct {
	Id         int            `gorm:"primary_key"`
	Username   string         `gorm:"not null"`
	Password   string         `gorm:"not null"`
	Status     int            `gorm:"not null"`
	Created_at string         `gorm:"not null"`
	Updated_at string         `gorm:"not null"`
	Deleted_at sql.NullString `gorm:"null"`
}

//按用户名查一个
func FindOneMember(username string) (Member, error) {
	var m Member
	err := global.DB.Where("username = ?", username).First(&m)
	if err.Error != nil {
		fmt.Printf("findone data failed err :%s\n", err.Error)
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

// 插入一条数据
func InsertOneMember(username string, password string) (err error) {
	m := Member{Username: username, Password: password, Status: 1, Deleted_at: sql.NullString{Valid: false}, Created_at: time.Now().Format(global.TimeFormat), Updated_at: time.Now().Format(global.TimeFormat)}
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

// 更新一个用户名和密码
func UpdateOneMember(username, password string) {
	ret, err := FindOneMember(username)
	if err != nil {
		fmt.Printf("find failed err:%s\n", err)
	}
	res := global.DB.Model(&ret).Select("username", "password", "updated_at").Updates(Member{Username: username, Password: password, Updated_at: time.Now().Format(global.TimeFormat)})
	if res.Error != nil {
		fmt.Printf("save failed err:%s\n", res.Error)
	}
	fmt.Printf("update success rows:%d\n", res.RowsAffected)
}

// 删除一个用户需要用户名
func DelOneMember(username string) {
	res, err := FindOneMember(username)
	if err != nil {
		fmt.Printf("find failed err:%s\n", err)
		return
	}
	ret := global.DB.Delete(&Member{}, res.Id)
	fmt.Printf("update success rows:%d\n", ret.RowsAffected)
}
