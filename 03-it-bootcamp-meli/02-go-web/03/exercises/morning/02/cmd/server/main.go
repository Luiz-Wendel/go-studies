/*
	Exercício 2 - Gerar o método DELETE

		É necessário implementar uma funcionalidade para excluir uma entidade.
		Para isso, é necessário seguir os seguintes passos:
			1. Gere um método DELETE para excluir a entidade com base no ID.
			2. Se não existir, retorne um erro 404..
*/

package main

import (
	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/03/exercises/morning/02/cmd/server/controllers"
	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/03/exercises/morning/02/internal/users"
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
