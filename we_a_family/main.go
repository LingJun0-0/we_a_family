package main

import "fmt"

func main() {

	//连接数据库
	err := InitDB()
	if err != nil {
		fmt.Printf("connection mysql db failed:%s", err)
	}
	fmt.Println("connection mysql db success")
	//main函数结束前数据库连接关闭
	defer Close()

	router := getGinApp()
	//router := get2()
	//router := gin.Default()

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
