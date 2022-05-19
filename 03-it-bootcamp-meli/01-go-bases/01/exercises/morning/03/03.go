package main

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
