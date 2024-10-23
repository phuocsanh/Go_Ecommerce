package initialize

import (
	"go_ecommerce/global"

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

// v1 := r.Group("v1/2024")
// {
// 	v1.GET("/getUser",c.NewUserController().GetUserByID)
// }
// v2 := r.Group("v2/2024")
// {
// 	v2.GET("/ping",Pong)
// }
return r
}

 