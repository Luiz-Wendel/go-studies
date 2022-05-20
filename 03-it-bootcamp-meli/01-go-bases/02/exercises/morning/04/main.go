/*
	Exercício 4 - Cálculo de estatísticas

		Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio de suas notas.
		Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo, máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi indicado na função anterior.

		Exemplo:

		```
			const (
			minimum = "minimum"
			average = "average"
			maximum = "maximum"
			)

			...
			minFunc, err := operation(minimum)
			averageFunc, err := operation(average)
			maxFunc, err := operation(maximum)

			...
			minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
			averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
			maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		```
*/

package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	numbers := []int{2, 3, 3, 4, 1, 2, 4, 5}
	operations := []string{minimum, average, maximum}

	for _, operation := range operations {
		if function, err := getOperationFunction(operation); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(operation, "-", function(numbers...))
		}
	}
}

func getOperationFunction(operation string) (func(numbers ...int) int, error) {
	functions := map[string]func(numbbers ...int) int{
		minimum: getMinimum,
		average: getAverage,
		maximum: getMaximum,
	}

	if function, ok := functions[operation]; ok {
		return function, nil
	}

	return nil, errors.New("operation not recognized")
}

func getMinimum(numbers ...int) int {
	min := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
	}

	return min
}

func getAverage(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum / len(numbers)
}

func getMaximum(numbers ...int) int {
	max := numbers[0]

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}
