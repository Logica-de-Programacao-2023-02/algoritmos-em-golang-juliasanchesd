package main

import "fmt"

func main() {
	var salario_inicial float64
	fmt.Print("Qual seu salário atual? ")
	fmt.Scan(&salario_inicial)

	var salario_final float64
	salario_final = salario_inicial * (1.15)
	fmt.Printf("Seu salário com um aumento de 15%% é: %.2f", salario_final)

}
