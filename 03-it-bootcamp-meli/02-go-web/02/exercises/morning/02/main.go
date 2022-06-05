/*
	Exercício 2 - Validação de campo

		As validações dos campos devem ser implementadas no momento do envio do pedido, para isso devem ser seguidos os seguintes passos:
			1. Todos os campos enviados na solicitação devem ser validados, todos os campos são obrigatórios
			2. Caso algum campo não esteja completo, um código de erro 400 deve ser retornado com a mensagem “campo %s é obrigatório”.
				 (Em %s deve ir o nome do campo que não está completo).
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
	FirstName    string `json:"firstName" binding:"required"`
	LastName     string `json:"lastName" binding:"required"`
	Email        string `json:"email" binding:"required"`
	CreationDate string `json:"creationDate" binding:"required"`
	Age          uint   `json:"age" binding:"required"`
	Height       uint   `json:"height" binding:"required"`
	Active       bool   `json:"active" binding:"required"`
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

func getUserValue(u user, key string) string {
	reflectedUser := reflect.ValueOf(u)

	return fmt.Sprintf("%v", reflectedUser.FieldByName(strings.Title(key)))
}

func filterUsers(payload []user, filters map[string]string) []user {
	var filteredPayload []user

	for _, u := range payload {
		isValid := true

		for filter, value := range filters {
			userValue := getUserValue(u, filter)

			isValid = userValue == value
		}

		if isValid {
			filteredPayload = append(filteredPayload, u)
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
		validationString := " Error:Field validation"
		requiredString := "required"
		errorMessage := err.Error()

		if strings.Contains(errorMessage, requiredString) {
			errorIndex := strings.Index(errorMessage, validationString)

			field := errorMessage[11 : errorIndex-1]

			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("field \"%s\" is required", strings.ToLower(field)),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errorMessage,
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
