package main

import "fmt"

func main() {
	var dias float64
	fmt.Print("Quantos dias você trabalhou? ")
	fmt.Scan(&dias)

	var diaria float64
	fmt.Print("Quanto você ganhou por dia? ")
	fmt.Scan(&diaria)

	salario := dias * diaria
	fmt.Print("Seu salário foi de: ", salario)

}
