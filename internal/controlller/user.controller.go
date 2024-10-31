package controlller

import (
	"fmt"
	"go_ecommerce/internal/service"
	"go_ecommerce/internal/vo"
	"go_ecommerce/pkg/response"
	"log"

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
	log.Println("Email params ")
	var params vo.UserRegistratorRequest
	fmt.Print("Email params %s",params.Email)
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrResponse(c,response.ErrParamInvalid, err.Error())
	}
	

	result:=uc.userService.Register(params.Email,params.Purpose)
	response.SuccessResponse(c,result,nil)
}
