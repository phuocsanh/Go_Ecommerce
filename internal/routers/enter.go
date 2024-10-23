package routers

import (
	"go_ecommerce/internal/routers/manage"
	"go_ecommerce/internal/routers/user"
)


type RouterGroup struct{
	User user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)