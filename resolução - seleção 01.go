package main

import "fmt"

func main() {
	var num01 int
	fmt.Print("Informe o primeiro número inteiro: ")
	fmt.Scan(&num01)

	var num02 int
	fmt.Print("Informe o segundo número inteiro: ")
	fmt.Scan(&num02)

	var maior int
	if num01 > num02 {
		maior = num01
	} else {
		maior = num02
	}
	fmt.Print("O maior número entre os dois é: ", maior)

}
