package initialize

import (
	"github.com/gin-gonic/gin"
	"ping/router"
	"net/http"
)

// 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS("static", http.Dir("static")) // 为文件提供静态地址
	Router.Use(CrosHandler())
	//路由组
	ApiGroup := Router.Group("/api/")
	router.InitDashBoardRouter(ApiGroup)                //报表相关路由
	return Router
}

func CrosHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*") 
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
 
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
 
		//处理请求
		context.Next()
	}
}