package main

import "fmt"

func main() {

	var firstNumber int = 20
	var secondNumber int = 30

	fmt.Println("Vrijednosti su ", firstNumber, " i ", secondNumber)
	firstNumber, secondNumber = secondNumber, firstNumber

	fmt.Println("Vrijednosti su ", firstNumber, " i ", secondNumber)
}
