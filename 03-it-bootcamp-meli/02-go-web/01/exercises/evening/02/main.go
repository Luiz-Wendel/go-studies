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
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

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

var fields = []string{
	"id",
	"firstName",
	"lastName",
	"email",
	"age",
	"height",
	"active",
	"creationDate",
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
			userValue := fmt.Sprintf("%v", reflectedUser.FieldByName(strings.Title(filter)))

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

func getUserById(payload []user, id string) (user, error) {
	for _, user := range payload {
		if user.Id == id {
			return user, nil
		}
	}

	return user{}, errors.New("user not found")
}

func GetById(c *gin.Context) {
	content := getFileContent("./03-it-bootcamp-meli/02-go-web/01/exercises/morning/01/users.json")

	var payload []user

	serializeJson(content, &payload)

	id := c.Param("id")

	user, err := getUserById(payload, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func main() {
	router := gin.Default()

	router.GET("/users", GetAll)
	router.GET("/users/:id", GetById)

	router.Run()
}
