package main

import (
	"we_a_family/we_a_family/core"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
)

func main() {
	core.InitConf()
	//初始化连接数据库
	global.DB = core.InitGorm()
	//数据结构models建数据库表并映射外键关系
	global.DB.AutoMigrate(&Models.Member{}, &Models.Tag{})
	global.DB.AutoMigrate(&Models.TagMember{})

}
