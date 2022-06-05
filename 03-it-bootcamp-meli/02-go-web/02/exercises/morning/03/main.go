/*
	Exercício 3 - Validar Token

		Para adicionar segurança à aplicação, o pedido deve ser enviado com um token, para isso devem ser seguidos os seguintes passos:
		1. No momento do envio da solicitação, deve ser validado que um token é enviado
		2. Esse token deve ser validado em nosso código (o token pode ser codificado permanentemente).
		3. Caso o token enviado não esteja correto, devemos retornar um erro 401 e uma mensagem que "você não tem permissão para fazer a solicitação solicitada".
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
	token := c.GetHeader("token")

	if token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "permission denied",
		})
		return
	}

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
