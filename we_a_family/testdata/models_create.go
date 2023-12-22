package main

import (
	"we_a_family/we_a_family/core"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
)

// 根据模型创建数据表
func main() {
	core.InitConf()
	//初始化连接数据库
	global.DB = core.InitGorm()
	//数据结构models建数据库表并映射外键关系
	global.DB.AutoMigrate(&Models.Member{}, &Models.Picture{}, &Models.Tag{})

}
