package main

import "fmt"

// Sku Entity
type Sku struct {
	name     string
	price    float64
	discount float64
	quantity int
}

type product struct {
	id          int
	name        string
	description string
	skus        []Sku
}

func (p product) calculateDiscount() float64 {

	total := 0.0
	for _, item := range p.skus {
		total += (item.price * (1 - item.discount)) * float64(item.quantity)
	}
	return total
}

func main() {

	skus := []Sku{
		Sku{name: "sku", price: 10, discount: 0.50, quantity: 3},
		Sku{name: "sku", price: 20, discount: 0.50, quantity: 10},
	}
	myproduct := product{1, "product", "description", skus}

	fmt.Println(myproduct.calculateDiscount())
}
