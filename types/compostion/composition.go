package main

import "fmt"

type esportive interface {
	turnTurbo()
}

type luxury interface {
	doGoal()
}

type esportiveLuxury interface {
	esportive
	luxury
}

type bmw7 struct{}

func (b bmw7) turnTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) doGoal() {
	fmt.Println("Balize...")
}

func main() {
	var car esportiveLuxury = bmw7{}
	car.turnTurbo()
	car.doGoal()
}
