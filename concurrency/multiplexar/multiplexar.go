package main

import (
	"fmt"
	"time"

	"github.com/aferlim/html"
)

func forward(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

// multiplexar - misturar (mensagem) num canal

func getTogether(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)

	go forward(entry1, c)
	go forward(entry2, c)

	return c
}

func main() {

	now := time.Now()

	c := getTogether(
		html.Titulo("https://www.nasa.gov", "https://www.google.com"),
		html.Titulo("https://mundocaixa.webpremios.com.br", "https://www.youtube.com"),
	)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println(time.Since(now))
}
