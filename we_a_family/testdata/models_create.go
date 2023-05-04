package main

import (
	"fmt"
	"time"
	"we_a_family/we_a_family/core"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
)

type User struct {
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}

func testMapping() {
	// 查询用户表中的数据
	var mebbers []Models.Member
	result := global.DB.Find(&mebbers)
	if result.Error != nil {
		// 处理错误
		fmt.Println("1", result.Error)
	}

	// 处理查询结果
	for _, user := range mebbers {
		createdAt, err := time.ParseInLocation("2006-01-02 15:04:05", user.CreatedAt.Format("2006-01-02 15:04:05"), time.Local)
		if err != nil {
			// 处理错误
			fmt.Println("2", result.Error)

		}
		user.CreatedAt = createdAt
		fmt.Println("user.CreatedAt", user.CreatedAt)

		updatedAt, err := time.ParseInLocation("2006-01-02 15:04:05", user.UpdatedAt.Format("2006-01-02 15:04:05"), time.Local)
		if err != nil {
			// 处理错误
			fmt.Println("3", result.Error)

		}
		user.UpdatedAt = updatedAt
		fmt.Println("user.UpdatedAt", user.UpdatedAt)

		// 处理每个用户
	}
}

func main() {
	core.InitConf()
	//初始化连接数据库
	global.DB = core.InitGorm()
	testMapping()
	////数据结构models建数据库表并映射外键关系
	//global.DB.AutoMigrate(&Models.Member{}, &Models.Picture{}, &Models.Tag{}, &Models.Perm{})

}
