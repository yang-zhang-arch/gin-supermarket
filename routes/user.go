package routes

import (
	"WebFull/controller"
	"WebFull/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(eng *gin.Engine) *gin.Engine {
	eng.Use(middleware.CSRFMiddleware()) // 跨域请求处理中间件
	eng.POST("/api/auth/register", controller.Register)
	eng.POST("/api/auth/login", controller.Login)
	eng.GET("/api/auth/info", controller.Info)

	return eng

}
