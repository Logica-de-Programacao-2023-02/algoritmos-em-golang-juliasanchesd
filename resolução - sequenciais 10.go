package main

import "fmt"

func main() {
	var peso_quilos float64
	fmt.Print("Qual o seu peso em kg? ")
	fmt.Scan(&peso_quilos)

	peso_libras := peso_quilos * 2.20462
	fmt.Printf("Seu peso em libras é: %.2f ", peso_libras)
}
