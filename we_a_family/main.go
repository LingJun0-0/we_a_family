package main

func main() {

	router := getGinApp()
	//router := get2()
	//router := gin.Default()

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
