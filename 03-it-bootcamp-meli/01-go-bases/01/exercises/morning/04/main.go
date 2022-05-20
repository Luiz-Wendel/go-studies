/*
	Exercício 4 - Tipos de dados
		Um estudante de programação tentou fazer declarações de variáveis com seus respectivos tipos em Go mas teve várias dúvidas ao fazê-lo. A partir disso, ele nos deu seu código e pediu a ajuda de um desenvolvedor experiente que pode:
			● Revisar o código e realizar as devidas correções.

			```
				var sobrenome string = "Silva"
				var idade int = "25"
				boolean := "false";
				var salario string = 4585.90
				var nome string = "Fellipe"
			```
*/

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
