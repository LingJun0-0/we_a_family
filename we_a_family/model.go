package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
)

func gen_title_by_num() gin.H {

	// 模拟查库
	num := rand.Intn(10)
	var value string
	if num <= 1 {
		value = "垃圾小胖"
	} else {
		value = "垃圾毛强"
	}
	return gin.H{
		"title": value,
	}
}
