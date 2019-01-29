package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f4(p1, p2 string, p4 int) string {
	return fmt.Sprintf("F2: %s %s %d\n", p1, p2, p4)
}

func f5(p1 string) (string, string) {
	return "1", p1
}

func main() {
	f1()
	fmt.Println(f4("1", "2", 3))

	r1, r2 := f5("10")
	fmt.Println(r1, r2)

	fmt.Println(f5("10"))
}
