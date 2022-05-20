package main

/*
	Exercício 2 - Clima
		Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e pressão atmosférica de diferentes lugares.
			1. Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
			2. Imprima os valores no console.
			3. Quais tipos de dados serão atribuídos a essas variáveis?
*/

import "fmt"

func main() {
	temperature := 12.8
	umidity := 38
	atmosfericPressure := 30.09

	fmt.Printf("Temperature\t\t%v\t%T\n", temperature, temperature)
	fmt.Printf("Umidity\t\t\t%v\t%T\n", umidity, umidity)
	fmt.Printf("Atmosferic pressure\t%v\t%T\n", atmosfericPressure, atmosfericPressure)
}
