package perm

import (
	"pages/component"
	"pages/httpd/handler/perm"
)

func Routes() {
	pRoute := component.Router.Group("/permission")
	{
		pRoute.POST("", perm.AddPerm)
	}
}
