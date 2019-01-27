package main

import "fmt"

// has no ternary operator

func getResult(note float64) string {
	if note >= 6 {
		return "Approved"
	}
	return "Reproved"

	// return note > 6 ? "Approved" : "Reproved"
}

func main() {
	fmt.Println(getResult(6.2))
}
