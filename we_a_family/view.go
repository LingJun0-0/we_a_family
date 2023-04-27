package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		context.AsciiJSON(http.StatusOK, data)
		//c.JSON(http.StatusOK, data)
	})

	// html load
	router.LoadHTMLGlob("D:\\Program Files\\GoPath\\src\\we_a_family\\templates\\*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gen_title_by_num())
	})

	return
}
