package main

import "fmt"

func conceptNote(n float64) string {

	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 {
		return "B"
	} else if n >= 5 {
		return "C"
	} else if n >= 3 {

		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(conceptNote(0))
}
