package user

import (
	"pages/component"
	"pages/httpd/handler/user"
)

func Routes() {
	uRoute := component.Router.Group("/user")
	{
		uRoute.POST("", user.Create(""))
	}
}
