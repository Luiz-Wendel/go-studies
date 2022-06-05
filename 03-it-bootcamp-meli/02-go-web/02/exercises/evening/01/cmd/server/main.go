/*
	Exercício 1 - Gerar pacote interno

		A estrutura do projeto deve ser separada e como primeiro passo gerando o pacote interno, todas as funcionalidades que não dependem de pacotes externos devem estar no pacote interno.

		Dentro do pacote devem estar as camadas:
			1. Service, deve conter a lógica da nossa aplicação.
				a. O arquivo service.go deve ser criado.
				b. A interface Service com todos os seus métodos deve ser gerada.
				c. A estrutura de serviço que contém o repositório deve ser gerada.
				d. Deve ser gerada uma função que retorne o Serviço.
				e. Todos os métodos correspondentes às operações a serem executadas (GetAll, Create, etc.) devem ser implementados.

			2. Repository, você deve ter acesso à variável armazenada na memória.
				a. O arquivo repository.go deve ser criado
				b. A estrutura da entidade deve ser criada
				c. Variáveis globais devem ser criadas para armazenar as entidades
				d. A interface Repository deve ser gerada com todos os seus métodos
				e. A estrutura do repositório deve ser gerada
				f. Deve ser gerada uma função que retorne o Repositório
				g. Todos os métodos correspondentes às operações a serem executadas (GetAll, Store, etc.) devem ser implementados.
*/

package main

import (
	"fmt"

	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/02/exercises/evening/01/internal/users"
)

func main() {
	repository := users.NewRepository()
	service := users.NewService(repository)

	usrs, err := service.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(usrs)

	var john, jane users.User

	john, err = service.Store("John", "Doe", "john.doe@gmail.com", "01/01/2001", 27, 180, true)
	if err != nil {
		fmt.Println(err)
	}

	usrs, err = service.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(usrs)

	jane, err = service.Store("Jane", "Doe", "jane@doe.com", "01/01/2010", 72, 168, false)
	if err != nil {
		fmt.Println(err)
	}

	usrs, err = service.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(usrs)

	service.Update(john.Id, "John", "Doe", "john@doe.com", "01/01/2001", 27, 182, true)

	usrs, err = service.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(usrs)

	service.Delete(jane.Id)

	usrs, err = service.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(usrs)
}
