package main

import "fmt"

func main() {
	var altura float64
	fmt.Print("Qual sua altura em metros? ")
	fmt.Scan(&altura)

	var peso float64
	fmt.Print("Qual seu peso em kg? ")
	fmt.Scan(&peso)

	IMC := peso / (altura * altura)
	fmt.Printf("Seu IMC é: %.2f", IMC)

}
