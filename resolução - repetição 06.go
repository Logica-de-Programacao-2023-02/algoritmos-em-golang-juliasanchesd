package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número para a tabuada: ")
	fmt.Scan(&numero)

	fmt.Printf("Tabuada de multiplicação para %d:\n", numero)

	var i int
	var resultado int
	for i = 1; i <= 10; i++ {
		resultado = numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
		fmt.Println()
	}

}
