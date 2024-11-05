package controllers

import (
	"net/http"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := []models.Task{
		{ID: 1, Name: "PHP", Status: false},
		{ID: 2, Name: "MYSQL", Status: false},
	}
	c.JSON(http.StatusOK, tasks)
}
