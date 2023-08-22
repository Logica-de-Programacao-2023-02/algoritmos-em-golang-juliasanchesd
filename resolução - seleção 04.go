package main

import "fmt"

func main() {
	var altura float64
	fmt.Print("Qual sua altura em metros? ")
	fmt.Scan(&altura)

	var peso float64
	fmt.Print("Qual seu peso? ")
	fmt.Scan(&peso)

	IMC := peso / (altura * altura)

	var resultado string
	if IMC < 18.5 {
		resultado = ("abaixo do peso ideal")
	}
	if IMC >= 18.5 && IMC <= 24.9 {
		resultado = ("dentro do peso ideal")
	}
	if IMC >= 25 && IMC <= 29.9 {
		resultado = ("acima do peso ideal")
	}
	if IMC > 30 {
		resultado = ("obeso")
	}
	fmt.Println()
	fmt.Println("Você está:", resultado)
}
