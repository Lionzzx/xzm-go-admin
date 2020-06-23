package initialize

import (
	"ginEssential/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")

	router.InitUserRouter(ApiGroup)

	return Router
}
