package product

import (
	"pages/model/product"

	"github.com/gin-gonic/gin"
)

// Get get product
func Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		results, _ := product.FindAll()
		// results := []string{"test", "go"}
		c.JSON(200, results)
	}
}
