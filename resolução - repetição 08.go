package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&numero)

	fmt.Printf("Os divisores de %d são:\n", numero)
	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			fmt.Println(i)
		}
	}
}
