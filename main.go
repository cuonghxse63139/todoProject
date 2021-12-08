package main

import (
	"github.com/gin-gonic/gin"
	"todoProject/controllers"
	"todoProject/middleware"
)

func main() {
	router := gin.Default()
	router.Use(gin.CustomRecovery(middleware.Recover()))
	userController := controllers.UserControllerStruct{}
	todoController := controllers.TodoControllerStruct{}

	todoRoutes := router.Group("/todo")
	{
		todoRoutes.GET("/", middleware.CheckToken(), todoController.GetTodo)
		todoRoutes.POST("/add", middleware.CheckToken(), todoController.InsertTodo)
		todoRoutes.PUT("/update", middleware.CheckToken(), todoController.UpdateTodo)
		todoRoutes.DELETE("/:id", middleware.CheckToken(), todoController.DeleteTodo)
	}

	router.POST("/login", userController.Login)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
