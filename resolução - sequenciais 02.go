package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Informe um número inteiro ")
	fmt.Scan(&numero)

	var dobro int
	dobro = numero * 2

	var triplo int
	triplo = numero * 3

	var quadruplo int
	quadruplo = numero * 4

	fmt.Println("O dobro desse número é: ", dobro)
	fmt.Println("O triplo desse número é: ", triplo)
	fmt.Println("O quadruplo desse número é: ", quadruplo)

}
