package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Informe um número inteiro: ")
	fmt.Scan(&numero)

	var resultado string
	if numero%2 == 0 {
		resultado = ("par")
	} else {
		resultado = ("ímpar")
	}
	fmt.Println("O número é:", resultado)

}
