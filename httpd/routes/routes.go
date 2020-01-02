package routes

import (
	"pages/component"
	nonauth "pages/httpd/routes/misc"
	"pages/httpd/routes/perm"
	"pages/httpd/routes/product"
	"pages/httpd/routes/user"
	"pages/middleware/auth"
)

// Routes list all routes
func Routes() {
	nonauth.Routes()
	component.Router.Use(auth.JWT())
	component.Router.Use(auth.Casbin())
	product.Routes()
	user.Routes()
	perm.Routes()
}
