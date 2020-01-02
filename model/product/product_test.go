package product

import "testing"

func TestAdd(t *testing.T) {
	products := New()
	products.Add(Product{"Test", "Go"})
	if len(products.Products) != 1 {
		t.Errorf("Product was not added")
	}
}

func TestGetAll(t *testing.T) {
	products := New()
	products.Add(Product{"Test", "Go"})
	if len(products.GetAll()) != 1 {
		t.Errorf("Get all products failed")
	}
}
