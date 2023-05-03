package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"we_a_family/we_a_family/global"
	"we_a_family/we_a_family/utils"
)

type Member struct {
	Id         int              `gorm:"primary_key"json:"id"`
	Username   int              `gorm:"unique_index"`
	Password   string           `gorm:"type:varchar(255);not null"`
	Status     utils.StatusCode `json:"status"`
	Deleted    bool             `json:"deleted"`
	Created_at time.Time        `json:"created_At"`
	Updated_at time.Time        `json:"updated_At"`
	Tag        []Tag            `gorm:"foreignKey:MemberId"`
	TagMember  []TagMember      `gorm:"foreignKey:MemberId"`
}

//第一次查用户按用户名
func LoginFindMember(username string) (Member, error) {
	var m Member
	err := global.DB.Where("username = ?", username).First(&m)
	if err.Error != nil {
		fmt.Printf("findone member failed err :%s\n", err.Error)
	}
	return m, err.Error
}

//按用户id查用户信息
func FindOneMemberById(id int) (Member, error) {
	var m Member
	err := global.DB.Where("id = ?", id).First(&m)
	if err.Error != nil {
		fmt.Printf("findone member failed err :%s\n", err.Error)
	}
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
func InsertOneMember(username int, password string) (err error) {
	m := Member{Username: username, Password: password, Status: utils.MemberStatusCode1, Deleted: false, Created_at: time.Now(), Updated_at: time.Now()}
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
func UpdateOneMemberById(id, username int, password string, delete bool) error {
	res := global.DB.Model(&Member{}).Where("id = ?", id).Select("username", "password", "delete", "updated_at").Updates(Member{Username: username, Password: password, Deleted: delete, Updated_at: time.Now()})
	if res.Error != nil {
		fmt.Printf("save failed err:%s\n", res.Error)
		return res.Error
	}
	fmt.Printf("update success rows:%d\n", res.RowsAffected)
	return nil
}

// 按id删除一个用户
func DelOneMemberById(id int) error {
	ret := global.DB.Model(&Member{}).Where("id = ?", id).Select("deleted").Updates(Member{Deleted: true})
	if ret.Error != nil {
		fmt.Printf("save failed err:%s\n", ret.Error)
		return ret.Error
	}
	fmt.Printf("update success rows:%d\n", ret.RowsAffected)
	return nil
}
