package main

import "fmt"

type Address struct {
	grad  string
	ulica string
}

type Osoba struct {
	ime    string
	godine int
	adresa Address
}

func (o Osoba) Ispis() {
	fmt.Println(o.ime, " , ", o.godine, " godine ", " , zivi u ", o.adresa.grad, " , ", o.adresa.ulica)
}

func main() {

	adr := Address{grad: "Mostar", ulica: "Franjevacka"}

	osb := Osoba{ime: "Blaz", godine: 22, adresa: adr}

	osb.Ispis()

}
