package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Subtraction =>", a-b)
	fmt.Println("Division =>", a/b)
	fmt.Println("Multiplication =>", a*b)
	fmt.Println("Module =>", a%b)

	// bitwise
	fmt.Println("AND =>", a&b) // 11 & 10 = 10
	fmt.Println("OR =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// operations using math

	fmt.Println("Major => ", math.Max(c, d))
	fmt.Println("Minor => ", math.Min(c, d))
	fmt.Println("Exponation => ", math.Pow(c, d))

	fmt.Println("Major => ", math.Max(float64(a), float64(b)))

}