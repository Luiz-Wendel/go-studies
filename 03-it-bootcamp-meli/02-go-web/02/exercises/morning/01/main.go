/*
	Exercício 1 - Criar Entidade

		A funcionalidade para criar a entidade deve ser implementada. Se isso acontecer, os seguintes passos devem ser seguidos:
			1. Crie um endpoint por meio de POST que receba a entidade.
			2. Você deve ter um array da entidade na memória (no nível global), no qual todas as requisições que são feitas devem ser salvas.
			3. No momento de fazer a solicitação, o ID deve ser gerado. Para gerar o ID, devemos procurar o ID do último registro gerado, incrementá-lo em 1 e atribuí-lo ao nosso novo registro (sem ter uma variável global do último ID).
*/

package main

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

var users []user

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

func filterUsers(payload []user, filters map[string]string) []user {
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
	filters := getFilters(c)

	fmt.Println(filters)

	filteredUsers := filterUsers(users, filters)

	c.JSON(http.StatusOK, gin.H{
		"data": filteredUsers,
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
	id := c.Param("id")

	user, err := getUserById(users, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func CreateUser(c *gin.Context) {
	var u user

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	u.Id = uuid.NewString()

	users = append(users, u)

	c.JSON(http.StatusOK, gin.H{
		"data": u,
	})
}

func main() {
	router := gin.Default()

	usersGroup := router.Group("/api/users")
	{
		usersGroup.GET("/", GetAll)
		usersGroup.GET("/:id", GetById)
		usersGroup.POST("/", CreateUser)
	}

	router.Run()
}
