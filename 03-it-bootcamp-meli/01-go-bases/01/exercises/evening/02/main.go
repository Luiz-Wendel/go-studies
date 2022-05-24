/*
	Exercício 2 - Empréstimo

		Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
		Para isso, possui certas regras para saber a qual cliente pode ser concedido.
		Apenas concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano de atividade.
		Dentro dos empréstimos que concede, não cobra juros para quem tem um salário superior a US$ 100.000.
		É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem de acordo com cada caso.

		Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.
*/

package main

import "fmt"

type person struct {
	age, workingMonths, salary uint
	isEmployed                 bool
}

var (
	minimalAge    uint = 22
	isEmployed    bool = true
	workingMonths uint = 12
	salary        uint = 100000
)

func main() {
	p1 := person{age: 20, workingMonths: workingMonths, salary: salary, isEmployed: isEmployed}
	p2 := person{age: minimalAge, workingMonths: 11, salary: salary, isEmployed: isEmployed}
	p3 := person{age: minimalAge, workingMonths: workingMonths, salary: salary, isEmployed: false}
	p4 := person{age: minimalAge, workingMonths: workingMonths, salary: 6927, isEmployed: isEmployed}
	p5 := person{age: minimalAge, workingMonths: workingMonths, salary: salary, isEmployed: isEmployed}

	loans := []string{checkLoan(p1), checkLoan(p2), checkLoan(p3), checkLoan(p4), checkLoan(p5)}

	for _, loan := range loans {
		fmt.Println(loan)
	}
}

func checkLoan(p person) string {
	if p.age < minimalAge || p.isEmployed != isEmployed || p.workingMonths < workingMonths {
		return "Loan denied"
	}

	if p.salary < salary {
		return "Loan approved with interest"
	}

	return "Loan approved free of interest"
}
