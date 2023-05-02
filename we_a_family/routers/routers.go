package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"we_a_family/we_a_family/models"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()
	Router.POST("/insertone/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		Models.InsertOneMember(username, password)
		m, _ := Models.FindOneMember(username)
		c.JSON(http.StatusOK, m.Id)
	})
	Router.GET("/findone/:username", func(c *gin.Context) {
		username := c.Param("username")
		m, _ := Models.FindOneMember(username)
		c.JSON(http.StatusOK, m.Id)
	})
	Router.GET("/findoneUseid/:id", func(c *gin.Context) {
		id := c.Param("id")
		id1, _ := strconv.Atoi(id)
		m, _ := Models.FindOneMemberUseId(id1)
		c.JSON(http.StatusOK, m.Username)
	})
	Router.GET("/findall", func(c *gin.Context) {
		m := Models.FindsAllMember()
		c.JSON(http.StatusOK, m[1].Id)
	})
	Router.PATCH("/deleteone/:id/:username", func(c *gin.Context) {
		id := c.Param("id")
		username := c.Param("username")
		id1, _ := strconv.Atoi(id)
		Models.DelOneMember(id1)
		m, _ := Models.FindOneMember(username)
		c.JSON(http.StatusOK, m.Deleted)
	})
	Router.PATCH("/updatedone/:id/:username/:password/:delete", func(c *gin.Context) {
		id := c.Param("id")
		id1, _ := strconv.Atoi(id)
		username := c.Param("username")
		password := c.Param("password")
		delete := c.Param("delete")
		delete1, _ := strconv.ParseBool(delete)

		Models.UpdateOneMember(id1, username, password, delete1)
		m, _ := Models.FindOneMember(username)
		c.JSON(http.StatusOK, m.Password)
	})
	return Router
}
