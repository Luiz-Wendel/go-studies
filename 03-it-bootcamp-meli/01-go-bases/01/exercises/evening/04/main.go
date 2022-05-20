/*
	Exercício 4 - Quantos anos tem...
		Um funcionário de uma empresa deseja saber o nome e a idade de um de seus funcionários.
		De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

		```
			var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
		```

		Por outro lado, você também precisa:
			- Saiba quantos de seus funcionários têm mais de 21 anos.
			- Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
			- Excluir Pedro do mapa.
*/

package main

import "fmt"

func main() {
	employees := map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("Benjamin -", employees["Benjamin"], "year(s) old")

	var ofAge uint

	for _, employee := range employees {
		if employee > 21 {
			ofAge += 1
		}
	}

	fmt.Println("Employees of age:", ofAge)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}
