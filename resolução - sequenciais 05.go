package main

import "fmt"

func main() {
	var idade_anos int
	fmt.Print("Qual sua idade? ")
	fmt.Scan(&idade_anos)

	idade_dias := idade_anos * 365
	fmt.Print("Sua idade em dias é: ", idade_dias)
}
