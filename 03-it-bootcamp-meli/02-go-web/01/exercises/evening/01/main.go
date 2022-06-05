/*
	Exercício 1 - Vamos filtrar nosso endpoint

	Dependendo do tema escolhido, precisamos adicionar filtros ao nosso endpoint, ele deve ser capaz de filtrar todos os campos.
		1. Dentro do manipulador de endpoint, recebi os valores para filtrar do contexto.
		2. Em seguida, ele gera a lógica do filtro para nossa matriz.
		3. Retorne a matriz filtrada por meio do endpoint.
*/

package main

import (
	"encoding/json"
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

func main() {
	router := gin.Default()

	router.GET("/users", GetAll)

	router.Run()
}
