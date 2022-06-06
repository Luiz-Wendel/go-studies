/*
	Exercício 3 - Gerar o método PATCH

		É necessário implementar uma funcionalidade que modifique parcialmente a entidade, apenas 2 campos devem ser modificados:
			- Se Produtos for selecionado, os campos de nome e preço.
			- Se Usuários for selecionado, os campos sobrenome e idade.
			- Se Transações foi selecionado, o código da transação e os campos de valor.
		Para conseguir isso, é necessário seguir os seguintes passos:
			1. Gere um método PATCH para modificar parcialmente a entidade, modificando apenas
			2 campos (por opção).
			2. Do Path envie o ID da entidade a ser modificada.
			3. Se não existir, retorne um erro 404.
			4. Valide os 2 campos para enviados.
*/

package main

import (
	"log"

	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/03/exercises/evening/02/cmd/server/controllers"
	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/03/exercises/evening/02/internal/users"
	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/03/exercises/evening/02/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error when loading the .env file")
	}

	db := store.New(store.FileType, "./users.json")
	repository := users.NewRepository(db)
	service := users.NewService(repository)
	userController := controllers.NewUserController(service)

	router := gin.Default()

	usersRoutes := router.Group("/api/users")
	{
		usersRoutes.GET("/", userController.GetAll())
		usersRoutes.POST("/", userController.Store())
		usersRoutes.PUT("/:id", userController.Update())
		usersRoutes.PATCH("/:id", userController.PartialUpdate())
		usersRoutes.DELETE("/:id", userController.Delete())
	}

	router.Run()
}
