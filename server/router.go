package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shaoji-org/shaoji-appreciate/api"
	"github.com/shaoji-org/shaoji-appreciate/middleware"
	"os"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	// r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		//// 用户登录
		//v1.POST("user/register", api.UserRegister)
		//
		//// 用户登录
		//v1.POST("user/login", api.UserLogin)
		//
		//// 需要登录保护的
		//authed := v1.Group("/")
		//authed.Use(middleware.AuthRequired())
		//{
		//	// User Routing
		//	authed.GET("user/me", api.UserMe)
		//	authed.DELETE("user/logout", api.UserLogout)
		//}

		// 视频操作
		v1.POST("shaoji", api.CreateShaoJi)
		v1.GET("shaoji/:id", api.ShowShaoJi)
		v1.GET("shaojis", api.ListShaoJi)
		v1.PUT("shaoji/:id", api.UpdateShaoJi)
		v1.DELETE("shaoji/:id", api.DeleteShaoJi)
		// 排行榜
		v1.GET("rank/daily", api.DailyRank)
		// 其他
		v1.POST("upload/token", api.UploadToken)
	}

	// swagger文档
	// 游览器打开 http://localhost:3000/swagger/index.html
	//r.StaticFile("/swagger.json", "./swagger/swagger.json")
	//r.Static("/swagger", "./swagger/dist")

	return r
}
