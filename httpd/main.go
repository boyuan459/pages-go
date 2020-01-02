package main

import (
	"pages/component"
	"pages/httpd/routes"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	defer component.DB.Close()
	routes.Routes()
	component.Router.Run()
}
