package main

import (
	"fmt"
	"time"

	"github.com/aferlim/html"
)

func theFaster(url, url2, url3 string) string {
	c1 := html.Titulo(url)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3

	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {

	now := time.Now()

	winner := theFaster(
		"https://www.youtube.com",
		"https://www.nasa.gov",
		"https://www.google.com",
	)

	fmt.Println(winner)
	fmt.Println(time.Since(now))
}
