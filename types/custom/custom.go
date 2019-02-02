package main

import "fmt"

type note float64

func (n note) between(begin, end float64) bool {
	return float64(n) >= begin && float64(n) <= end
}

// Concept T
func Concept(n note) string {
	if n.between(9.0, 10.0) {
		return "A"
	} else if n.between(7.0, 8.99) {
		return "B"
	} else if n.between(5.0, 6.99) {
		return "C"
	} else {
		return "D"
	}
}

func main() {

	// n := 9.0
	fmt.Println(Concept(9.0))
	fmt.Println(Concept(8.0))
	fmt.Println(Concept(6.0))
	fmt.Println(Concept(4.5))
}
