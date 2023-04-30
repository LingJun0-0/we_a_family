package main

import (
	"we_a_family/core"
	"we_a_family/global"
)

func main() {
	/*方法执行顺序
	1，先读取配置文件
	2，初始化各个组件
	*/

	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.Initlogger()
	//初始化连接数据库
	global.DB = core.InitGorm()

}
