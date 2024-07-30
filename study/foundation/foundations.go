package main

import "fmt"

// Declarando variável constante
const a = "Hello, World!"

// Criando um tipo
type ID int

// Declarando variaveis
var (
	b bool    = true
	c int     = 10
	d string  = "João"
	e float64 = 1.2
	f ID      = 1
)

// Método main
func main() {
	// Criando uma variável do tipo Array
	var meuArray [3]int
	// Atribuindo os valores no array
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3
	// Rodando o for com os valores do array
	for i, v := range meuArray {
		fmt.Printf("O valor do index %d é %d\n", i, v)
	}
	a := "X" // atribuindo uma variável local
	b = false
	println(a)
	//println(f)
	fmt.Printf("O tipo de F é %T\n", f)
}
