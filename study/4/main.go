package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Joao"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de F é %T\n", f)
}
