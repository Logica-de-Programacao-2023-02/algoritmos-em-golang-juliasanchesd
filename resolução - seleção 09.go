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


	if num01 <= num02 && num02 <= num03 && num01 <= num03 {
	        fmt.Printf("%d < %d < %d" , num01 , num02 , num03) 
	} if num02 <= num01 && num02 <= num03 && num01 <= num03
	        fmt.Printf("%d < %d < %d" , num02 , num01 , num03)
        } if num01 <= num02 && num03 <= num02 && num01 <= num03 {
	        fmt.Printf("%d < %d < %d" , num01 , num03 , num02 
        } if num02 <= num01 && num03 <= num01 && num02 <= num03 {
		fmt.Printf("%d < %d < %d" , num02 , num03 , num01
	} if num03 <= num01 && num01 <= num02 && num03 <= num02 {
		fmt.Printf("%d < %d < %d" , num03 , num01 , num02)
	} if num03 <= num02 && num02 <= num01 && num03 <= num01 {
		fmt.Printf("%d < %d < %d" , num03 , num02 , num01)
		
}
