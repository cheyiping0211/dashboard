package router

import (
	"github.com/gin-gonic/gin"
	api "ping/api"
)

func InitDashBoardRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("dashboard")
	{
		ApiRouter.GET("buildModule", api.BuildModules)     //打包
		ApiRouter.GET("getModule", api.GetModules)  //获取当前项目的报表配置
		ApiRouter.PUT("updateModule", api.UpdateModules) 	//修改当前项目的报表配置
	}
}
