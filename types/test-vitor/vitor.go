package main

import "fmt"

type Brinquedo interface {
	brincar()
	falar()
	andar()
	pular()
	correr()
	rir()
}

type HomemAranha struct {
}

type ThannosGlove struct {
}

type Wow struct {
}

type Dino struct {
}

func (b HomemAranha) brincar() {
	fmt.Println("homem aranha não brinca")
}

func (b HomemAranha) falar() {
	fmt.Println("homem aranha fala")
}

func (b HomemAranha) andar() {
	fmt.Println("homem aranha não anda")
}

func (b HomemAranha) pular() {
	fmt.Println("homem aranha pula")
}

func (b HomemAranha) correr() {
	fmt.Println("homem aranha não corre")
}

func (b HomemAranha) rir() {
	fmt.Println("homem aranha ri")
}

func (b ThannosGlove) brincar() {
	fmt.Println("luva do Thannos não brinca")
}

func (b ThannosGlove) falar() {
	fmt.Println("luva do Thannos não fala")
}

func (b ThannosGlove) andar() {
	fmt.Println("luva do Thannos não anda")
}

func (b ThannosGlove) pular() {
	fmt.Println("luva do Thannos não pula")
}

func (b ThannosGlove) correr() {
	fmt.Println("luva do Thannos não corre")
}

func (b ThannosGlove) rir() {
	fmt.Println("luva do Thannos não ri")
}

func main() {
	var brinquedo Brinquedo = ThannosGlove{}

	brinquedo.andar()
	brinquedo.brincar()
	brinquedo.correr()
	brinquedo.falar()
	brinquedo.pular()
	brinquedo.rir()
}
