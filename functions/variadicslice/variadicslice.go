package main

import "fmt"

func printApproved(approvedL ...string) {
	fmt.Println("Approveds List")

	for i, approved := range approvedL {

		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main() {
	approvedL := []string{"Manoel", "Maria", "Rafael"}

	printApproved(approvedL...)
}
