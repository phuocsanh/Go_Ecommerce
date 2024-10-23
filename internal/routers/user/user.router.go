package user

import (
	"go_ecommerce/internal/controlller"
	"go_ecommerce/internal/repo"
	"go_ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{

}
func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup){
// public router
ur:=repo.NewUserRepository()
us:= service.NewUserService(ur)
userHandlerNonDependency:=controlller.NewUserController(us)

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("register", userHandlerNonDependency.Register)
		userRouterPublic.POST("/otp")
	}
	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("get_info")
	}
}