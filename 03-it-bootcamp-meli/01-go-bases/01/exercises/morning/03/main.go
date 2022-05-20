package main

/*
	Exercício 3 - Declaração de variáveis

		Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é declarar diferentes variáveis.
		Ajude o professor com as seguintes questões:
			1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
			2. Corrigir as incorrectas.

			```
				var 1nome string
				var sobrenome string
				var int idade
				1sobrenome := 6
				var licenca_para_dirigir = true
				var estatura da pessoa int
				quantidadeDeFilhos := 2
			```
*/

import "fmt"

func main() {
	var nome string
	var sobrenome string
	var idade int
	sobrenome = "Doe"
	licencaParaDirigir := true
	var estaturaDaPessoa int
	quantidadeDeFilhos := 2

	fmt.Printf("%v\t%T\n", nome, nome)
	fmt.Printf("%v\t%T\n", sobrenome, sobrenome)
	fmt.Printf("%v\t%T\n", idade, idade)
	fmt.Printf("%v\t%T\n", licencaParaDirigir, licencaParaDirigir)
	fmt.Printf("%v\t%T\n", estaturaDaPessoa, estaturaDaPessoa)
	fmt.Printf("%v\t%T\n", quantidadeDeFilhos, quantidadeDeFilhos)
}
