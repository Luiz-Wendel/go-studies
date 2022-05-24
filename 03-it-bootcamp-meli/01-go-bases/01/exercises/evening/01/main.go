/*
	Exercício 1 - Letras de uma palavra

		A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma das letras separadamente para soletrá-la. Para isso terão que:
			1. Crie uma aplicação que tenha uma variável com a palavra e imprima o número de letras que ela contém.
			2. Em seguida, imprimi cada uma das letras.
*/

package main

import "fmt"

func main() {
	word := "paralelepipedo"

	fmt.Println("Word length:", len(word))

	for _, char := range word {
		fmt.Println(string(char))
	}
}
