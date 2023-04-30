package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"we_a_family/we_a_family/models"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()
	Router.POST("/insertone", func(c *gin.Context) {
		Models.InsertOneMember("lizhen", "lz")
		m, _ := Models.FindOneMember("lizhen")
		c.JSON(http.StatusOK, m.Id)
	})
	Router.POST("/findone", func(c *gin.Context) {
		m, _ := Models.FindOneMember("genghao")
		c.JSON(http.StatusOK, m.Id)
	})
	Router.POST("/findall", func(c *gin.Context) {
		m := Models.FindsAllMember()
		c.JSON(http.StatusOK, m[0].Id)
	})
	Router.POST("/deleteone", func(c *gin.Context) {
		Models.DelOneMember(1)
		m, _ := Models.FindOneMember("genghao")
		c.JSON(http.StatusOK, m.Deleted)
	})
	Router.POST("/updatedone", func(c *gin.Context) {
		Models.UpdateOneMember(1, "genghao", "genghao", false)
		m, _ := Models.FindOneMember("genghao")
		c.JSON(http.StatusOK, m.Password)
	})
	return Router
}
