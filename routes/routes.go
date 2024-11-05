package routes

import (
	"todo-list/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers)
	r.GET("/tasks", controllers.GetTasks)
}
