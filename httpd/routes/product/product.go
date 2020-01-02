package product

import (
	"pages/component"
	"pages/httpd/handler/product"
)

func Routes() {
	pRoute := component.Router.Group("/product")
	{
		pRoute.GET("", product.Get())
	}
}
