package main

import (
	"fmt"
	"strconv"
)

func main() {
	// função auxiliar para formatar números
	formatNum := func(num float64) string {
		return strconv.FormatFloat(num, 'f', 2, 64)
	}

	for {
		var num1, num2 float64
		var opcao string
		var resultado string
		var statusOperacao = true

		fmt.Println("===== Calculadora =====")
		fmt.Print("Digite o primeiro número: ")
		fmt.Scan(&num1)
		fmt.Print("Digite o segundo número: ")
		fmt.Scan(&num2)
		fmt.Print(`
			Escolha uma opção:
			1. Soma
			2. Subtração
			3. Multiplicação
			4. Divisão
			5. Maior número
			0. Sair
			Opção: 
		`)
		fmt.Scan(&opcao)

		if opcao == "0" {
			fmt.Println("Encerrando o programa...")
			break
		}

		switch opcao {
		case "1":
			resultado = "Soma: " + formatNum(num1+num2)
		case "2":
			resultado = "Subtração: " + formatNum(num1-num2)
		case "3":
			resultado = "Multiplicação: " + formatNum(num1*num2)
		case "4":
			if num2 != 0 {
				resultado = "Divisão: " + formatNum(num1/num2)
			} else {
				statusOperacao = false
				fmt.Println("Erro: Divisão por zero não é permitida!")
			}
		case "5":
			if num1 > num2 {
				resultado = formatNum(num1) + " é maior."
			} else if num2 > num1 {
				resultado = formatNum(num2) + " é maior."
			} else {
				resultado = "Os números são iguais"
			}
		default:
			statusOperacao = false
			fmt.Println("Erro: Opção inválida!")
		}

		if statusOperacao {
			fmt.Println("Opção: ", opcao)
			fmt.Print("Resultado: ", resultado)
		}
	}

}
