package main

import "fmt"

type Course struct {
	name string
}

func main() {
	var thing interface{}

	fmt.Println(thing)

	thing = 3
	fmt.Println(thing)

	thing = true
	fmt.Println(thing)

	type dinamic interface{}

	var thing2 dinamic = "Opa"
	fmt.Println(thing2)

	thing2 = true
	fmt.Println(thing2)

	thing2 = Course{"Vitola"}
	fmt.Println(thing2)

}
