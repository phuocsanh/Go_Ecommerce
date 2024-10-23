package routers

import (
	// c "go_ecommerce/internal/controlller"
	"go_ecommerce/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
r := gin.Default()
r.Use(middlewares.AuthenMiddleware())

// v1 := r.Group("v1/2024")
{
	// v1.GET("/getUser",c.NewUserController().GetUserByID)
}
// v2 := r.Group("v2/2024")
// {
// 	v2.GET("/ping",Pong)
// }
return r
}

