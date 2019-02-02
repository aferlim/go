package main

import "fmt"

type Sportive interface {
	turnOnTurbo()
}

type Ferrari struct {
	model    string
	turbo    bool
	velocity int
}

func (f *Ferrari) turnOnTurbo() {
	f.turbo = true

}

func main() {
	car1 := Ferrari{"F40", false, 80}
	car1.turnOnTurbo()

	var car2 Sportive = &Ferrari{"F40", false, 80}
	car2.turnOnTurbo()

	fmt.Println(car1, car2)
}
