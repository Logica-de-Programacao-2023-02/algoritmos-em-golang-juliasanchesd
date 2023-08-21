package main

import "fmt"

func main() {
	var preço float64
	fmt.Print("Qual o preço do produto? ")
	fmt.Scan(&preço)

	desconto := preço * 0.9
	fmt.Printf("O valor do produto com 10%% de desconto é: %.2f ", desconto)
}
