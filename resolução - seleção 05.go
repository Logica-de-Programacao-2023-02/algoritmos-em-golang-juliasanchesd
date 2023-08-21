package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Informe um número inteiro ")
	fmt.Scan(&numero)

	var resultado string
	if numero%3 == 0 && numero%5 == 0 {
		resultado = ("é múltiplo de 3 e 5")
	} else {
		resultado = ("não é múltiplo de 3 e 5")
	}

	fmt.Println("O número:", resultado)
}
