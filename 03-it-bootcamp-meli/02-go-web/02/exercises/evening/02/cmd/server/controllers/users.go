package controllers

import (
	"fmt"
	"net/http"

	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/02/exercises/evening/02/internal/users"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service users.Service
}

func NewUserController(service users.Service) *UserController {
	return &UserController{service}
}

func (c *UserController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": u,
		})
	}
}

type request struct {
	FirstName    string `json:"firstName" binding:"required"`
	LastName     string `json:"lastName" binding:"required"`
	Email        string `json:"email" binding:"required"`
	CreationDate string `json:"creationDate" binding:"required"`
	Age          uint   `json:"age" binding:"required"`
	Height       uint   `json:"height" binding:"required"`
	Active       *bool  `json:"active" binding:"required"`
}

func (c *UserController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "INVALIDINPUT-001",
				"error":   "invalid inputs",
				"message": err.Error(),
			})
			return
		}

		u, err := c.service.Store(req.FirstName, req.LastName, req.Email, req.CreationDate, req.Age, req.Height, *req.Active)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":  "NOTFOUND-002",
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": u,
		})
	}
}

func (c *UserController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "INVALIDINPUT-001",
				"error":   "invalid inputs",
				"message": err.Error(),
			})
			return
		}

		u, err := c.service.Update(id, req.FirstName, req.LastName, req.Email, req.CreationDate, req.Age, req.Height, *req.Active)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":  "NOTFOUND-002",
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": u,
		})
	}
}

func (c *UserController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":  "NOTFOUND-002",
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": fmt.Sprintf("user with id: \"%s\" removed", id),
		})
	}
}
