package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	// cuidado ...
	fmt.Println("test " + string(97))

	fmt.Println("test " + strconv.Itoa(97))

	fmt.Println("test " + strconv.Itoa(notaFinal))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("1")

	if b {
		fmt.Println("true")
	}
}
