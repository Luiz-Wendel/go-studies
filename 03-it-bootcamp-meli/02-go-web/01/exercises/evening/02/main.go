/*
	Exercício 2 - Get one endpoint

		Gere um novo endpoint que nos permita buscar um único resultado do array de temas.
		Usando parâmetros de caminho o endpoint deve ser /theme/:id (lembre-se que o tema sempre tem que ser plural).
		Uma vez que o id é recebido, ele retorna a posição correspondente.

			1. Gere uma nova rota.
			2. Gera um manipulador para a rota criada.
			3. Dentro do manipulador, procure o item que você precisa.
			4. Retorna o item de acordo com o id.

		Se você não encontrou nenhum elemento com esse id retorne como código de resposta 404.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id, FirstName, LastName, Email, CreationDate string
	Age, Height                                  uint
	Active                                       bool
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

var fields = []string{
	"Id",
	"FirstName",
	"LastName",
	"Email",
	"Age",
	"Height",
	"Active",
	"CreationDate",
}

func getFilters(c *gin.Context) map[string]string {
	filters := make(map[string]string)

	for _, field := range fields {
		if value := c.Query(field); value != "" {
			filters[field] = value
		}
	}

	return filters
}

func filterPayload(payload []user, filters map[string]string) []user {
	var filteredPayload []user

	for _, user := range payload {
		isValid := true
		reflectedUser := reflect.ValueOf(user)

		for filter, value := range filters {
			userValue := fmt.Sprintf("%v", reflectedUser.FieldByName(filter))

			isValid = userValue == value
		}

		if isValid {
			filteredPayload = append(filteredPayload, user)
		}
	}

	return filteredPayload
}

func GetAll(c *gin.Context) {
	content := getFileContent("./03-it-bootcamp-meli/02-go-web/01/exercises/morning/01/users.json")

	var payload []user

	serializeJson(content, &payload)

	filters := getFilters(c)

	fmt.Println(filters)

	filteredPayload := filterPayload(payload, filters)

	c.JSON(http.StatusOK, gin.H{
		"message": filteredPayload,
	})
}

func GetById(c *gin.Context) {
	content := getFileContent("./03-it-bootcamp-meli/02-go-web/01/exercises/morning/01/users.json")

	var payload []user

	serializeJson(content, &payload)

	id := c.Param("id")

	for _, user := range payload {
		if user.Id == id {
			c.JSON(http.StatusOK, gin.H{
				"message": user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "user not found",
	})
}

func main() {
	router := gin.Default()

	router.GET("/users", GetAll)
	router.GET("/users/:id", GetById)

	router.Run()
}
