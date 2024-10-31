package user

import (
	"go_ecommerce/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{

}
func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup){
// public router

	/*
	Nếu không dùng dependency injection
	ur:=repo.NewUserRepository()
	us:= service.NewUserService(ur)
	userHandlerNonDependency:=controlller.NewUserController(us)
	*/

	//Dùng dependency injection by wire
	userController, _:= wire.InitUserRouterHandler()
	userRouterPublic := Router.Group("/user")
	{
		/*
		Nếu không dùng dependency injection
		userRouterPublic.POST("register", userHandlerNonDependency.Register)
		*/

		//Dùng dependency injection by wire
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}
	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("get_info")
	}
}