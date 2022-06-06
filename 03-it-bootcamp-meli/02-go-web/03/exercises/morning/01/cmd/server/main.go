/*
	Exercício 1 - Gerar o método PUT

		É solicitado a implementação de uma funcionalidade que modifique completamente uma entidade.
		Para isso, é necessário seguir os seguintes passos:
			1. Gere um método PUT para modificar toda a entidade.
			2. No Path envie o ID da entidade a ser modificada.
			3. Se não existir, retorne um erro 404..
			4. Realize todas as validações (todos os campos são obrigatórios).
*/

package main

import (
	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/03/exercises/morning/01/cmd/server/controllers"
	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/03/exercises/morning/01/internal/users"
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
	}

	router.Run()
}
