package main

import "fmt"

func multi(a, b int) int {
	return a * b
}

func exec(function func(int, int) int, p1, p2 int) int {
	return function(p1, p2)
}

func main() {
	result := exec(multi, 3, 4)
	fmt.Println(result)
}
