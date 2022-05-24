/*
	Exercício 3 - Calcular salário

		Uma empresa marítima precisa calcular o salário de seus funcionários com base no número de horas trabalhadas por mês e na categoria do funcionário.
		Se a categoria for C, seu salário é de R$1.000 por hora
		Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h mensais
		Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h mensais

		Calcule o salário dos funcionários conforme as informações abaixo:
			- Funcionário de categoria C: 162h mensais
			- Funcionário de categoria B: 176h mensais
			- Funcionário de categoria A: 172h mensais
*/

package main

import "fmt"

type employee struct {
	category    string
	workingTime uint
}

type payment struct {
	salary uint
	bonus  float64
}

func main() {
	e1 := employee{category: "C", workingTime: 162}
	e2 := employee{category: "B", workingTime: 176}
	e3 := employee{category: "A", workingTime: 172}

	fmt.Println(e1.category, e1.workingTime, getPayment(e1))
	fmt.Println(e2.category, e2.workingTime, getPayment(e2))
	fmt.Println(e3.category, e3.workingTime, getPayment(e3))
}

func getPayment(e employee) float64 {
	payments := map[string]payment{
		"C": {salary: 1000, bonus: 1.0},
		"B": {salary: 1500, bonus: 1.2},
		"A": {salary: 3000, bonus: 1.5},
	}

	if e.workingTime > 160 {
		return float64(payments[e.category].salary) * float64(e.workingTime) * payments[e.category].bonus
	} else {
		return float64(payments[e.category].salary) * float64(e.workingTime)
	}
}
