package main

import "fmt"

type printable interface {
	toString() string
}

type Person struct {
	name    string
	surname string
}

type Product struct {
	name  string
	price float64
}

func (p Person) toString() string {
	return fmt.Sprintf("%s %s", p.name, p.surname)
}

func (p Product) toString() string {
	return fmt.Sprintf("%s - %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing printable = Person{"Andre", "Lima"}
	fmt.Println(thing.toString())
	print(thing)

	thing = Product{"sku", 9.9}
	fmt.Println(thing.toString())
	print(thing)

	thing2 := Product{"sku", 19.9}
	print(thing2)
}
