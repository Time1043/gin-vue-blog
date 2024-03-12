package routers

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	/*router.GET("", func(context *gin.Context) {
		context.String(200, "xxx")
	})  // 验证成功*/

	apiRouterGroup := router.Group("api")         // 路由分组 127.0.0.1:8080/api
	routerGroupApp := RouterGroup{apiRouterGroup} // 路由分层
	routerGroupApp.SettingsRouter()               // 系统配置api
	return router
}
