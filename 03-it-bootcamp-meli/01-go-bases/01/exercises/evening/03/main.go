/*
	Exercício 3 - A que mês corresponde?

		Faça uma aplicação que contenha uma variável com o número do mês.
		1. Dependendo do número, imprima o mês correspondente em texto.
		2. Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você escolheria e porquê?

		Ex: 7 de julho.
*/

package main

import "fmt"

func main() {
	numberedMonth := 7

	fmt.Println(getStringifiedMonth(numberedMonth))
}

func getStringifiedMonth(numberedMonth int) string {
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	stringifiedMonth := months[numberedMonth]

	if stringifiedMonth != "" {
		return stringifiedMonth
	} else {
		return "invalid value"
	}
}
