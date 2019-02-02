package main

import (
	"fmt"
	"strings"
)

// Person Format
type Person struct {
	name    string
	surname string
}

func (p Person) getFullName() string {
	return fmt.Sprintf("%s %s", p.name, p.surname)
}

func (p *Person) setFullName(fullname string) {

	names := strings.Split(fullname, " ")

	p.name = names[0]
	p.surname = names[1]
}

func main() {

	console := fmt.Println

	person := Person{name: "André", surname: "Lima"}

	console(person.getFullName())

	person.setFullName("Andre Ferreira Lima")

	console(person.getFullName())

}
