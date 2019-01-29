package main

import "fmt"

func printResult(note float64) {
	if note >= 7 {
		fmt.Println("Student grade approved with", note)
	} else {
		fmt.Println("Student not approved with", note)
	}
}

func main() {
	printResult(7.3)
	printResult(5.1)
}
