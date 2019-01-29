package main

import "fmt"

func change(p1, p2 int) (st1, nd2 int) {
	st1 = p2
	nd2 = p1
	return
}

func main() {

	fmt.Println(change(10, 50))
}
