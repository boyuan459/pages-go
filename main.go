package main

import (
	"fmt"

	toolbox "./"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := toolbox.ParseConfig("./config/app.json")
	if err != nil {
		fmt.Println("config error", err)
	}
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
