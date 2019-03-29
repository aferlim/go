package main

import "fmt"

func fatorial(n int) (int, error) {

	switch {
	case n < 0:
		return -1, fmt.Errorf("invalid number: %d", n)
	case n == 0:
		return 1, nil
	default:
		previousFatorial, _ := fatorial(n - 1)
		fmt.Println(n - 1)
		return previousFatorial, nil
	}
}

func main() {
	result, _ := fatorial(50)
	fmt.Println(result)

	// _, err := fatorial(-4)

	// if err != nil {
	// 	fmt.Println(err)
	// }
}
