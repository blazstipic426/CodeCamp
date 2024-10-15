package main

import "fmt"

type Pravokutnik struct {
	Sirina int
	Visina int
}

func (p Pravokutnik) Povrsina_Opseg() (int, int) {
	povrsina := p.Sirina * p.Visina
	Opseg := 2 * (p.Sirina + p.Visina)

	return povrsina, Opseg

}

func main() {

	prav := Pravokutnik{Sirina: 6, Visina: 4}

	povrsina, opseg := prav.Povrsina_Opseg()

	fmt.Println("Povrsina je", povrsina, " i opseg je ", opseg)

}
