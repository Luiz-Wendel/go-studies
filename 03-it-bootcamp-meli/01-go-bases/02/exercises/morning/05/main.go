/*
	Exercício 5 - Cálculo da quantidade de alimento

		Um abrigo de animais precisa descobrir quanta comida comprar para os animais de estimação.
		No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão é que haja muito mais animais para abrigar.
			1. Cães precisam de 10 kg de alimento
			2. Gatos 5 kg
			3. Hamster 250 gramas.
			4. Tarântula 150 gramas.

		Precisamos:
			1. Implementar uma função Animal que receba como parâmetro um valor do tipo texto com o animal especificado e que retorne uma função com uma mensagem (caso não exista o animal)
			2. Uma função para cada animal que calcule a quantidade de alimento com base na quantidade necessária do animal digitado.

		Exemplo:

		```
			const (
			dog = "dog"
			cat = "cat"
			)

			...
			animalDog, msg := Animal(dog)
			animalCat, msg := Animal(cat)

			...
			var amount float64
			amount+= animaldog(5)
			amount+= animalCat(8)
		```
*/

package main

import (
	"errors"
	"fmt"
)

const (
	DOG       = "dog"
	CAT       = "cat"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

func main() {
	animals := map[string]uint{
		DOG:       5,
		CAT:       9,
		HAMSTER:   22,
		TARANTULA: 9,
		"fox":     2,
	}

	amount := 0

	for animal, quantity := range animals {
		if function, err := Animal(animal); err != nil {
			fmt.Printf("%v - %v\n", animal, err)
		} else {
			amount += int(function(quantity))
		}
	}

	fmt.Println("Need to buy:", float64(amount)/1000, "kg of animal food")
}

func Animal(s string) (func(quantity uint) uint, error) {
	availableAnimals := map[string]func(q uint) uint{
		DOG:       getDogFoodQuantity,
		CAT:       getCatFoodQuantity,
		HAMSTER:   getHamsterFoodQuantity,
		TARANTULA: getTarantulaFoodQuantity,
	}

	if function, ok := availableAnimals[s]; ok {
		return function, nil
	}

	return nil, errors.New("animal not available")
}

func getDogFoodQuantity(q uint) uint {
	return q * 10000
}

func getCatFoodQuantity(q uint) uint {
	return q * 5000
}

func getHamsterFoodQuantity(q uint) uint {
	return q * 250
}

func getTarantulaFoodQuantity(q uint) uint {
	return q * 150
}
