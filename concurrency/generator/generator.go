package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go concurrency Patterns

// <-chan - canal somente leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title(.*?)>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[2]
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.nasa.gov", "https://www.google.com")
	t2 := titulo("https://www.grupoltm.com.br", "https://www.youtube.com")

	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)

}