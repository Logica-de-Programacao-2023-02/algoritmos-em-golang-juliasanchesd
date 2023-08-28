package main

import "fmt"

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

	if num01 <= num02 && num02 <= num03 && num01 <= num03 {
		fmt.Printf("%.1f < %.1f < %.1f", num01, num02, num03)
	} else if num02 <= num01 && num02 <= num03 && num01 <= num03 {
		fmt.Printf("%.1f < %.1f < %.1f", num02, num01, num03)
	} else if num01 <= num02 && num03 <= num02 && num01 <= num03 {
		fmt.Printf("%.1f < %.1f < %.1f", num01, num03, num02)
	} else if num02 <= num01 && num03 <= num01 && num02 <= num03 {
		fmt.Printf("%.1f < %.1f < %.1f", num02, num03, num01)
	} else if num03 <= num01 && num01 <= num02 && num03 <= num02 {
		fmt.Printf("%.1f < %.1f < %.1f", num03, num01, num02)
	} else if num03 <= num02 && num02 <= num01 && num03 <= num01 {
		fmt.Printf("%.1f < %.1f < %.1f", num03, num02, num01)
	}

}
