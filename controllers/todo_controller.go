package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"todoProject/config"
	"todoProject/dtos"
	"todoProject/entities"
	"todoProject/services"
)

type TodoController interface {
	GetTodo(c *gin.Context)
	InsertTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type TodoControllerStruct struct {
	service services.TodoService
}

func (controller *TodoControllerStruct) init() error {
	tempService := &services.TodoServiceStruct{}
	controller.service = tempService
	return controller.service.Init()
}

func (controller *TodoControllerStruct) GetTodo(c *gin.Context) {
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	result, err := controller.service.GetListTodoByUserId(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			handleNotFound(
				c,
				dtos.BadRequestResponse{
					ErrorMessage: "Not found todo with ID " + fmt.Sprintf("%d", userId),
				})
			return
		}
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}
	handleSuccess(c, result)
}

func (controller *TodoControllerStruct) InsertTodo(c *gin.Context) {
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}
	todo := &entities.Todo{}
	if err := c.BindJSON(todo); err != nil {
		log.Println(err)
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}
	todo.UserId = userId
	resultTodo, err := controller.service.InsertTodo(todo)

	if err != nil {
		log.Println(err)
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, resultTodo)
}

func (controller *TodoControllerStruct) UpdateTodo(c *gin.Context) {
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}
	todo := &entities.Todo{}
	if err := c.BindJSON(todo); err != nil {
		log.Println(err)
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}
	todo.UserId = userId
	resultTodo, err := controller.service.UpdateTodo(todo)

	if err != nil {
		log.Println(err)
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, resultTodo)
}

func (controller *TodoControllerStruct) DeleteTodo(c *gin.Context) {
	var rowEffect int64
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}
	todoId, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}
	rowEffect, err = controller.service.DeleteTodo(todoId)

	if err != nil {
		log.Println(err)
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, rowEffect)
}

func getUserId(c *gin.Context) (int64, error) {
	var userId int64
	var error error
	temp, isOk := c.Get(config.TOKEN_CURRENT_USER_ID)
	if isOk {
		userId, error = strconv.ParseInt(fmt.Sprint(temp), 10, 64)
		if error == nil {
			return userId, nil
		}
	}
	return 0, nil
}
