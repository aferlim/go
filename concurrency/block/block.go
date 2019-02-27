package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)

	c <- 1 // operacao bloqueante

	fmt.Println("Só depois que for lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go routine(c)

	fmt.Println(<-c) // operação bloqueante
	fmt.Println("foi lido")

	//fmt.Println(<-c) //deadlock

}
