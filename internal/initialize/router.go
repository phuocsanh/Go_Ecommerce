package initialize

import (
	"go_ecommerce/global"
	"go_ecommerce/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r  *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode((gin.DebugMode))
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode((gin.ReleaseMode))
		r =gin.New()
	}

	//middleware
	// r.Use() //Logger
	// r.Use() //cross
	// r.Use() //Limiter global
	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User
	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}
 return r
}

 