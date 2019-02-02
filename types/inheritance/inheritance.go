package main

import "fmt"

type car struct {
	name            string
	currentVelocity int
}

type ferrari struct {
	car
	turboOn bool
}

func main() {
	ferr := ferrari{}

	ferr.name = "F40"
	ferr.currentVelocity = 0
	ferr.turboOn = true

	fmt.Println(ferr)
}
