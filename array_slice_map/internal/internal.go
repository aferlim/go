package main

import "fmt"

func main() {

	arr := make([]int, 10, 20)

	arr2 := append(arr, 1, 2, 3)

	fmt.Println(arr, arr2)

	arr[0] = 9

	fmt.Println(arr, arr2)
}
