/*
	Exercício 2 - Olá {nome}
		1. Crie dentro da pasta go-web um arquivo chamado main.go
		2. Crie um servidor web com Gin que retorne um JSON que tenha uma chave “mensagem” e diga Olá seguido do seu nome.
		3. Acesse o end-point para verificar se a resposta está correta.
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, John",
		})
	})

	router.Run()
}
