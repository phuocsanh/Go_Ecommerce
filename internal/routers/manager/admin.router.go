package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{

}
func (r *AdminRouter) InitAdminRouter(Router *gin.RouterGroup){

	// private router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")

	}

	// private router
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.POST("active_user")

	}
}