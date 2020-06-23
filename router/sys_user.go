package router

import (
	v1 "ginEssential/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
		UserRouter.POST("createUser", v1.CreateUser)
		UserRouter.POST("test", v1.Test)
	}
}

