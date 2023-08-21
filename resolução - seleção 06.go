package main

import "fmt"

func main() {
	var num01 int
	fmt.Print("Informe o primeiro número inteiro: ")
	fmt.Scan(&num01)

	var num02 int
	fmt.Print("Informe o segundo número inteiro: ")
	fmt.Scan(&num02)

	if num01 < 0 || num02 < 0 {
		resultado := num01 * num02
		fmt.Printf("A multiplicação dos dois números é: %d\n", resultado)
	}
	if num01 > 0 && num02 > 0 {
		resultado := num01 + num02
		fmt.Println("A soma dos dois números é:", resultado)
	}

}
