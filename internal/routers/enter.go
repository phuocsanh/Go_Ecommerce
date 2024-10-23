package routers

import (
	"go_ecommerce/internal/routers/manager"
	"go_ecommerce/internal/routers/user"
)


type RouterGroup struct{
	User user.UserRouterGroup
	Manager manager.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)