package routers

import (
	"github.com/gin-gonic/gin"
	"we_a_family/we_a_family/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouters() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//// 加载上两级目录的 layout.html 文件
	//gin.LoadHTMLFiles(filepath.Join("..", "layout.html"))
	//// 加载同一层级或下级目录的 home.html 文件
	//gin.LoadHTMLGlob(filepath.Join("*.html"))
	//router.LoadHTMLFiles("templates/user/login.html")
	router.LoadHTMLGlob("templates/user/*")

	//路由分组
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//路由分层
	//系统配置API
	routerGroupApp.SettingsRouter()
	//用户API
	routerGroupApp.MemberRouter()
	//图片API
	routerGroupApp.ImageRouter()

	return router
}
