package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha")

	fmt.Println("Nova")
	fmt.Println(" linha")

	x := 3.141516

	fmt.Println("o valor de x é ", x)

	xs := fmt.Sprint(x)

	fmt.Println("the value is " + xs)

	fmt.Printf("the x value is %.2f", x)

	a, b, c, d := 1, 1.9999, false, "Opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
