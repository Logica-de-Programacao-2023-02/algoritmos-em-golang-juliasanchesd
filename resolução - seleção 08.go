package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Informe um número entre 1 e 7: ")
	fmt.Scan(&numero)

	var resultado string
	if numero > 7 || numero < 1 {
		resultado = ("error")
		fmt.Print(resultado)
	}
	if numero == 1 {
		resultado = ("domingo")
		fmt.Println("O dia da semana correspondente ao seu número é:", resultado)
	}
	if numero == 2 {
		resultado = ("segunda-feira")
		fmt.Println("O dia da semana correspondente ao seu número é:", resultado)
	}
	if numero == 3 {
		resultado = ("terça-feira")
		fmt.Println("O dia da semana correspondente ao seu número é:", resultado)
	}
	if numero == 4 {
		resultado = ("quarta-feira")
		fmt.Println("O dia da semana correspondente ao seu número é:", resultado)
	}
	if numero == 5 {
		resultado = ("quinta-feira")
		fmt.Println("O dia da semana correspondente ao seu número é:", resultado)
	}
	if numero == 6 {
		resultado = ("sexta-feira")
		fmt.Println("O dia da semana correspondente ao seu número é:", resultado)
	}
	if numero == 7 {
		resultado = ("sábado")
		fmt.Println("O dia da semana correspondente ao seu número é:", resultado)
	}

}
