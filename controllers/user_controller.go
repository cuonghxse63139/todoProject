package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"todoProject/dtos"
	"todoProject/entities"
	"todoProject/services"
)

type UserController interface {
	Login(c *gin.Context)
}

type UserControllerStruct struct {
	service services.UserAccountService
}

func (controller *UserControllerStruct) init() error {
	tempService := &services.UserAccountServiceStruct{}
	controller.service = tempService
	return controller.service.Init()
}

func (controller *UserControllerStruct) Login(c *gin.Context) {
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}
	var user entities.User

	error := c.BindJSON(&user)

	if error != nil {
		log.Println(error)
		handleBadRequest(c, dtos.BadRequestResponse{
			ErrorMessage: error.Error(),
		})
	}

	resultUser, error := controller.service.Login(user)

	if error != nil {
		log.Println(error)
		handleError(c, dtos.BadRequestResponse{
			ErrorMessage: error.Error(),
		})
		return
	}

	handleSuccess(c, resultUser)
}
