/*
	Exercício 3 - Listar Entidade

		Já tendo criado e testado nossa API que nos recebe, geramos uma rota que retorna uma lista do tema escolhido.
			1. Dentro do "main.go", crie uma estrutura de acordo com o tema com os campos correspondentes.
			2. Crie um endpoint cujo caminho é /thematic (plural). Exemplo: “/products”
			3. Crie uma handler para o endpoint chamada "GetAll".
			4. Crie um slice da estrutura e retorne-a por meio de nosso endpoint.
*/

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	CreationDate string `json:"creationDate"`
	Age          uint   `json:"age"`
	Height       uint   `json:"height"`
	Active       bool   `json:"active"`
}

func getFileContent(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("error when opening file: ", err)
	}

	return content
}

func serializeJson(content []byte, variable interface{}) {
	err := json.Unmarshal(content, variable)
	if err != nil {
		log.Fatal("error while reading json: ", err)
	}
}

func GetAll(c *gin.Context) {
	content := getFileContent("./03-it-bootcamp-meli/02-go-web/01/exercises/morning/01/users.json")

	var payload []user

	serializeJson(content, &payload)

	c.JSON(http.StatusOK, gin.H{
		"message": payload,
	})
}

func main() {
	router := gin.Default()

	router.GET("/users", GetAll)

	router.Run()
}
