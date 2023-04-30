package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"we_a_family/we_a_family/core"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
)

func main() {

	//读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
	//初始化日志
	global.Log = core.Initlogger()
	global.Log.Warnln("123")
	global.Log.Error("123")
	global.Log.Infof("123")

	logrus.Warnln("123")
	logrus.Error("123")
	logrus.Infof("123")

	//连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
	m := Models.FindOneUser("genghao")
	fmt.Println(m.Username, m.Password)

	router := getGinApp()
	//router := get2()
	//router := gin.Default()

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
