package controlller

import (
	"go_ecommerce/internal/service"
	"go_ecommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

// type UserController struct {
// userService *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }
// func (uc *UserController) GetUserByID(c *gin.Context){

// 		// uid := c.Query("uid")
// 		// c.JSON(http.StatusOK, gin.H{ // map string
// 		// 	"massage":uc.userService.GetInfoUser(),
// 		// 	"uid": uid,
// 		// 	"users":[]string{"sang","Sang","Son"},
// 		// })
// 		// response.SuccessResponse(c , 20001,[]string{"Long", "Hưng", "Dũng"})
// 		response.ErrResponse(c, 20003,"")

// }

type UserController struct {
userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context){

	result:=uc.userService.Register("","")
	response.SuccessResponse(c,result,nil)
}