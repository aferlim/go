package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qtde int) {
	for index := 0; index < qtde; index++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", person, text, index+1)

	}
}

func main() {
	//go speak("Maria", "Pq vc não fala comigo", 500)
	//go speak("João", "Só posso falar depois com você!", 500)

	// time.Sleep(time.Second * 5)
	// fmt.Println("Fim!")

	go speak("Maria", "Entendi!!!", 10)
	//speak("João", "Parabéns!", 5)
}
