package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i
	*p++
	i += i
	i++

	// GO haven't arithimetic pointer
	//p++ // wrong

	fmt.Println(p, *p, i, &i, &p)
	//fmt.Println(i, &i)
}
