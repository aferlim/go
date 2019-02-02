package main

import "fmt"

type product struct {
	id          int
	name        string
	description string
	price       float64
	discount    float64
}

func (p product) calculateDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {

	myproduct := product{1, "sku", "description", 60, 0.10}

	fmt.Println(myproduct, myproduct.calculateDiscount())
}
