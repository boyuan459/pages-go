package product

import (
	"pages/component"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title    string
	Name     string
	Price    float64
	Quantity int
	Status   bool
}

func init() {
	component.DB.AutoMigrate(&Product{})
}

// FindAll find all products
func FindAll() ([]Product, error) {
	var products []Product
	component.DB.Find(&products)
	return products, nil
}
