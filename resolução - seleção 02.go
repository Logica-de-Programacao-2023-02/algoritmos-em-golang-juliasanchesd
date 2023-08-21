package main

import "fmt"

func main() {
	var num01 int
	fmt.Print("Informe o primeiro número: ")
	fmt.Scan(&num01)

	var num02 int
	fmt.Print("Informe o segundo número: ")
	fmt.Scan(&num02)

	var num03 int
	fmt.Print("Informe o terceiro número: ")
	fmt.Scan(&num03)

	var menor int
	if num01 < num02 && num01 < num03 {
		menor = num01
	}
	if num02 < num01 && num02 < num03 {
		menor = num02
	}
	if num03 < num02 && num03 < num01 {
		menor = num03
	}

	fmt.Println("O menor número entre os três é: ", menor)
}
