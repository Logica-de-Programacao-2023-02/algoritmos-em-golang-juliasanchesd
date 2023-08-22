package main

import (
	"fmt"
	"sort"
)

func main() {
	var num01 float64
	fmt.Print("Digite o primeiro número real: ")
	fmt.Scan(&num01)

	var num02 float64
	fmt.Print("Digite o segundo número real: ")
	fmt.Scan(&num02)

	var num03 float64
	fmt.Print("Digite o terceiro número real: ")
	fmt.Scan(&num03)

	numeros := []float64{num01, num02, num03}
	sort.Float64s(numeros)

	fmt.Println("Números em ordem crescente:", numeros)
}
