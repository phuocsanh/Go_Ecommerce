package initialize

import (
	"go_ecommerce/global"
	"go_ecommerce/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// r := gin.Default()
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// midderware
	// r.Use() // logging
	// r.Use() // cross
	// r.Use() // limiter global
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/check_status") // tracking monitor
	}
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	return r
}
