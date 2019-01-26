package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// tipos inteiros
	fmt.Println(1, 2, 1000)

	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (so positivos) .. uint8, uint16, uint32, uint64

	var b byte = 255 //uint8

	fmt.Println("o byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64

	fmt.Println("64 bits é", i1, reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)

	fmt.Println("O rune é", reflect.TypeOf(i2), i2)

	fmt.Println("Os tipos são", reflect.TypeOf('a'), reflect.TypeOf(9))

	// tipos reais (float32, float64)
	var x float32 = 49.99
	x2 := float32(49.99)

	fmt.Println("O tipo de x é", reflect.TypeOf(x), reflect.TypeOf(x2), x)
	fmt.Println("O tipo do literal é", reflect.TypeOf(49.99)) // float64

	// tipos boleanos
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))

	fmt.Println(!bo)

	//Strings
	s1 := "my name's Andre"

	fmt.Println("s1 type", reflect.TypeOf(s1))
	fmt.Println("s1 length", len(s1))

	// String com multiplas linhas

	s2 :=
		`hi
	my name's
	André`

	fmt.Println(s2)

	// Nao há tipo char

}
