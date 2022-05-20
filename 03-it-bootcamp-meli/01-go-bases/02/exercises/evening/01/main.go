/*
	Exercício 1 - Registro de estudantes

		Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os detalhes dos dados de cada um deles, conforme o exemplo abaixo:
			Nome: [Nome do aluno]
			Sobrenome: [Sobrenome do aluno]
			RG: [RG do aluno]
			Data de admissão: [Data de admissão do aluno]

		Os valores que estão entre colchetes devem ser substituídos pelos dados fornecidos pelos alunos.
		Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG, Data e que tenha um método de detalhamento.
*/

package main

import "fmt"

type Student struct {
	FirstName, LastName, RG, Date string
}

func (a Student) details() {
	fmt.Println("First name:", a.FirstName)
	fmt.Println("Last name:", a.LastName)
	fmt.Println("RG:", a.RG)
	fmt.Println("Admission date:", a.Date)
}

func main() {
	a1 := Student{FirstName: "Vini", LastName: "Gouveia", RG: "5555555", Date: "20/05/2022"}
	a2 := Student{FirstName: "José", LastName: "Moura", RG: "6666666", Date: "16/05/2022"}

	a1.details()
	fmt.Println()
	a2.details()
}
