package main

import "fmt"

func main() {
	var numero_01 int
	fmt.Print("Informe o primeiro número ")
	fmt.Scan(&numero_01)
	var numero_02 int
	fmt.Print("Informe o segundo número ")
	fmt.Scan(&numero_02)
	var numero_03 int
	fmt.Print("Informe o terceiro número ")
	fmt.Scan(&numero_03)

	var soma int
	soma = numero_01 + numero_02 + numero_03
	fmt.Println("A soma dos três números é: ", soma)

}
