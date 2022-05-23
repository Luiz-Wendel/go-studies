/*
	Exercício 1 - Impostos de salário

		Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o imposto de um salário.
		Temos a informação que um dos funcionários ganha um salário de R$50.000 e será descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse caso será descontado mais 10%.
*/

package main

import "fmt"

func main() {
	fmt.Println("Taxes (50000):", getTaxes(50000, 17))
	fmt.Println("Taxes (150000):", getTaxes(150000, 27))
}

func getTaxes(salary uint, tax uint) float64 {
	return float64(salary) * (float64(tax) / 100)
}
