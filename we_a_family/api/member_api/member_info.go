package member_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/utils"
)

func (MemberApi) MemberLoginInfoView(c *gin.Context) {
	username := c.Param("username")
	username1, _ := strconv.Atoi(username)
	password := c.Param("password")
	m, _ := Models.LoginFindMember(username1)
	if m.Password == password {
		utils.OkwithData(m, c)
	} else {
		utils.FailwithCode(utils.LoginError, c)
	}
}

func (MemberApi) MemberFindAll(c *gin.Context) {
	m := Models.FindsAllMember()
	for _, m1 := range m {
		if m1.Deleted {
			utils.FailwithCode(utils.LoginError, c)
		} else {
			utils.OkwithData(m1, c)
		}
	}
}

func (MemberApi) InsertMemberView(c *gin.Context) {
	username := c.Param("username")
	username1, _ := strconv.Atoi(username)
	password := c.Param("password")
	err := Models.InsertOneMember(username1, password)
	if err != nil {
		utils.FailwithCode(utils.LoginError, c)
	} else {
		member, _ := Models.LoginFindMember(username1)
		utils.OkwithData(member, c)

	}

}
