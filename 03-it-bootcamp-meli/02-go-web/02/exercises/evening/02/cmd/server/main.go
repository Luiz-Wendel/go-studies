/*
	Exercício 2 - Gerar pacote do server

		A estrutura do projeto deve ser separada, como segundo passo deve ser gerado o pacote do servidor onde serão adicionadas as funcionalidades do projeto que dependem de pacotes externos e o principal do programa.

		Dentro do pacote deve estar:
			1. O main do programa.
				a. O repositório, serviço e handler devem ser importados e injetados
				b. O roteador deve ser implementado para os diferentes endpoints
			2. O pacote handler com o controlador da entidade selecionada.
				a. A estrutura request deve ser gerada
				b. A estrutura do controller que tem o serviço como campo deve ser gerada
				c. A função que retorna o controlador deve ser gerada
				d. Todos os métodos correspondentes aos endpoints devem ser gerados
*/

package main

import (
	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/02/exercises/evening/02/cmd/server/controllers"
	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/02/exercises/evening/02/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repository := users.NewRepository()
	service := users.NewService(repository)
	userController := controllers.NewUserController(service)

	router := gin.Default()

	usersRoutes := router.Group("/api/users")
	{
		usersRoutes.GET("/", userController.GetAll())
		usersRoutes.POST("/", userController.Store())
		usersRoutes.PUT("/:id", userController.Update())
		usersRoutes.DELETE("/:id", userController.Delete())
	}

	router.Run()
}
