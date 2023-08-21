package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Qual o seu salário? ")
	fmt.Scan(&salario)

	var resultado float64
	if salario < 1000 {
		resultado = salario * 1.10
	}
	if salario >= 1000 {
		resultado = salario * 1.05
	}

	fmt.Println("O seu novo salário é:", resultado)
}
