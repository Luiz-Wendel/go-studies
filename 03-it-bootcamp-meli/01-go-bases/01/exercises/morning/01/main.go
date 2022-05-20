/*
	Exercício 1 - Imprimindo o nome na tela
		1. Crie uma aplicação que tenha uma variável “nome” e outra “idade”.
		2. Imprima no terminal o valor de cada variável.
*/

package main

import "fmt"

func main() {
	name := "John Doe"
	age := 42

	fmt.Printf("%v - %v\n", name, age)
}
