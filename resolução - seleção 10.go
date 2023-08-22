package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Qual sua idade? ")
	fmt.Scan(&idade)

	var resultado string
	if idade <= 9 {
		resultado = ("mirim")
	}
	if idade >= 10 && idade <= 13 {
		resultado = ("infantil")
	}
	if idade >= 14 && idade <= 17 {
		resultado = ("juvenil")
	}
	if idade > 18 {
		resultado = ("adulto")
	}

	fmt.Println("A sua classificação de acordo com sua idade é:", resultado)
}
