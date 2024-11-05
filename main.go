package main

import (
	"todo-list/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run("localhost:8000")
}
