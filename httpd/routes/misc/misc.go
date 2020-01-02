package misc

import (
	"pages/component"
	"pages/httpd/handler/user"
)

func Routes() {
	component.Router.POST("/login", user.Login())
	component.Router.POST("/signup", user.Create("user"))
}
