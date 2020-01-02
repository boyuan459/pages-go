package component

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgresql
)

var (
	// DB connection
	DB *gorm.DB
	// Router gin router
	Router *gin.Engine
)

// GetDB get db connection
func init() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dp password=pg_password sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	DB = db

	router := gin.Default()
	Router = router
}
