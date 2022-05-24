/*
	Exercício 2 - Calcular média

		Um colégio precisa calcular a média das notas (por aluno).
		Precisamos criar uma função na qual possamos passar N quantidade de números inteiros e devolva a média.

		Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro
*/

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	grades := []float64{5.7, 6.8, 5, 3.9, 8.7, 9.5}

	if mean, er := getMean(grades...); er != nil {
		fmt.Println(er)
	} else {
		fmt.Println("Your mean is:", mean)
	}
}

func getMean(grades ...float64) (float64, error) {
	var sum float64
	counter := 0

	for _, grade := range grades {
		if grade < 0 {
			return 0, errors.New("invalid grades")
		}

		sum += grade
		counter += 1
	}

	mean := sum / float64(counter)

	return math.Round(mean*10) / 10, nil
}
