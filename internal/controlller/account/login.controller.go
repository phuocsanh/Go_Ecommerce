package account

import (
	"go_ecommerce/internal/service"
	"go_ecommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

// management controller login user

var Login = new(cUserLogin)

type cUserLogin struct {

}

func (c *cUserLogin) Login(ctx *gin.Context){
	err:= service.UserLogin().Login(ctx)
	if err!=nil {
        response.ErrResponse(ctx, response.ErrParamInvalid, err.Error())
        return
    }
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}