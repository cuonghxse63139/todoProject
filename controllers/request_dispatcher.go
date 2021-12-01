package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoProject/dtos"
)

func handleBadRequest(c *gin.Context, errorResponse dtos.BadRequestResponse) {
	c.IndentedJSON(http.StatusBadRequest, errorResponse)
}

// handle success
func handleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func handleError(c *gin.Context, errorResponse dtos.BadRequestResponse) {
	c.JSON(http.StatusInternalServerError, errorResponse)
}

func handleNotFound(c *gin.Context, errorResponse dtos.BadRequestResponse) {
	c.JSON(http.StatusNotFound, errorResponse)
}
