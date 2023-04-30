package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Models "we_a_family/we_a_family/models"
)

//尽量只是把需求重导向具体的方法并返回结果到具体界面,不要处理数据
func register_url(router *gin.Engine) {

	// hello world
	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	// ascii json sample
	router.GET("/someJSON", func(context *gin.Context) {

		// 一部分数据相关工作放到了controller
		data := gen_decoder()
		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		//context.AsciiJSON(http.StatusOK, data)
		context.JSON(http.StatusOK, data)
	})

	// html load
	router.LoadHTMLGlob("templates/*.html")
	//router.LoadHTMLFiles("/templates/*")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", Models.Gen_title_by_num())
	})

	return
}
