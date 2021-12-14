package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"todoProject/dbConnention"
	"todoProject/dtos"
	"todoProject/entities"
	"todoProject/repository"
	"todoProject/services"
	"todoProject/unit"
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

func NewTodoController() *TodoControllerStruct {
	return &TodoControllerStruct{}
}

func (controller *TodoControllerStruct) init() error {
	db := dbConnention.NewDBConnecttion()
	repo, error := repository.NewTodoRepository(db)
	if error != nil {
		log.Println(error)
		return error
	}
	controller.service = services.NewTodoService(repo)
	return nil
}

func (controller *TodoControllerStruct) GetTodo(c *gin.Context) {
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}
	userId, err := unit.GetUserId(c)
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
	userId, err := unit.GetUserId(c)
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
	userId, err := unit.GetUserId(c)
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
