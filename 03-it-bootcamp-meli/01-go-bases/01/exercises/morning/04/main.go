package main

import "fmt"

func main() {
	var sobrenome string = "Silva"
	var idade int = 25
	boolean := false
	var salario float64 = 4585.90
	var nome string = "Fellipe"

	fmt.Printf("%v\t%T\n", sobrenome, sobrenome)
	fmt.Printf("%v\t%T\n", idade, idade)
	fmt.Printf("%v\t%T\n", boolean, boolean)
	fmt.Printf("%v\t%T\n", salario, salario)
	fmt.Printf("%v\t%T\n", nome, nome)
}
