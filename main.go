package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todoProject/controllers"
	"todoProject/entities"
)

var todoList []*entities.Todo

func main ()  {
	router := gin.Default()

	userController := controllers.UserControllerStruct{}

	router.GET("/todo", getTodos)
	router.POST("/add", addTodo)
	router.POST("/edit", editTodo)
	router.POST("/changeStatus", changeStatus)

	router.POST("/login", userController.Login)

	err := router.Run("localhost:8080")
	if err != nil {
		return 
	}
}

func getTodos(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK,todoList)
}

func addTodo(c *gin.Context)  {
	var newTodo *entities.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	newTodo.Id = getId()
	newTodo.Status = 1
	todoList = append(todoList, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func editTodo(c *gin.Context)  {
	var temp *entities.Todo
	if err := c.BindJSON(&temp); err != nil {
		return
	}
	if index := findById(temp.Id); index != 0 {
		todoList[index] = temp
		c.IndentedJSON(http.StatusOK, temp)
	}
}

func changeStatus(c *gin.Context)  {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	status, _ := strconv.Atoi(c.Param("status"))
	if index := findById(id); index != 0 {
		todoList[index].Status = status
		c.IndentedJSON(http.StatusOK, todoList[index])
		return
	}
}

func getId() int64 {
	var newId int64
	if len(todoList) == 0 {
		newId = 0
	}else {
		listSize := len(todoList)
		lastTodo := todoList[listSize - 1]
		newId = lastTodo.Id + 1
	}
	return newId
}

func findById(id int64) int {
	var temp int
	for idx, todo := range todoList {
		if id == todo.Id {
			temp = idx
		}
	}
	return temp
}

func login()  {
	
}
