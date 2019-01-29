package main

import (
	"fmt"
	"time"
)

func Type(i interface{}) string {

	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "we dont know"
	}
}

func main() {
	fmt.Println(Type(2.3))
	fmt.Println(Type(2))
	fmt.Println(Type("2.3"))
	fmt.Println(Type(func() {}))
	fmt.Println(Type(time.Now()))
}
