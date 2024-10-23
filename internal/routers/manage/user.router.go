package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{

}
func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup){

	// private router
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.POST("/active_user")

	}
}