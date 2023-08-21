package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Informe um número inteiro ")
	fmt.Scan(&numero)

	antecedor := numero - (1)

	sucessor := numero + 1

	fmt.Println("O antecedor desse número é:", antecedor)
	fmt.Println("O sucessor desse número é:", sucessor)
}
