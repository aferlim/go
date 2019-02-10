package main

import (
	"fmt"
	"time"
)

func speak(person string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", person, i)
		}
	}()

	return c
}

func getTogether(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-entry1:
				c <- s
			case s := <-entry2:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	c := getTogether(speak("João"), speak("Maria"))

	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
