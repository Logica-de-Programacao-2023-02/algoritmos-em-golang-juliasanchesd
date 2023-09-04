package main

import "fmt"

func main() {
	var soma, quantidade, x int
	x = -1
	for x != 0 {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)

		soma += x
		quantidade++
	}
	media := soma / quantidade
	fmt.Println("A média é ", media)
}
