package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64 = 10, 8

	var operador = "+"
	var titulo = "Calculadora"

	fmt.Println(titulo)
	if operador == "+" {
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	} else if operador == "-" {
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	} else if operador == "*" {
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	} else if operador == "/" {
		fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
	} else {
		fmt.Println("Operador inválido")
	}

}
