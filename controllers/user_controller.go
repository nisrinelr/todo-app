package controllers

import (
	"net/http"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{
		{ID: 1, Username: "Nisrine", Email: "lasfer@gmail.com"},
		{ID: 2, Username: "LASFER", Email: "nisrine@gmail.com"},
	}
	c.JSON(http.StatusOK, users)
}
