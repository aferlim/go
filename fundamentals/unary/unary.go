package main

import "fmt"

func main() {
	x := 1
	y := 2

	// no prefixed
	// postfix only
	x++

	fmt.Println(x)

	y--
	fmt.Println(y)

	// fmt.Println(x == y--)

}
