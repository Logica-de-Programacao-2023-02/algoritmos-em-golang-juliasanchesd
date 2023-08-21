package main

import "fmt"

func main() {
	var num1 float64
	fmt.Print("Informe o primeiro número ")
	fmt.Scan(&num1)

	var num2 float64
	fmt.Print("Informe o segundo número ")
	fmt.Scan(&num2)

	var num3 float64
	fmt.Print("Informe o terceiro número ")
	fmt.Scan(&num3)

	media := (num1*2 + num2*3 + num3*5) / 10
	fmt.Print("A média ponderada com os pesos 2, 3 e 5 respectivamente é: ", media)

}
