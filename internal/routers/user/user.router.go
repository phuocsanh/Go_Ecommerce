package user

import "github.com/gin-gonic/gin"

type UserRouter struct{

}
func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup){
// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("register")
		userRouterPublic.POST("/otp")

	}
	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("get_info")

	}
}