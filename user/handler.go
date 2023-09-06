package user

import (
	"bwastartup/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service Service
}

func NewUserHandler(service Service) *userHandler {
	return &userHandler{service}
}

func (h userHandler) RegisterUser(c *gin.Context) {
	var input RegisterUserInput
	// userInput := new(input)
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationsErrors(err)
		errorsMsg := gin.H{"errors": errors}
		response := helper.ApiResponse("Register account failed!", http.StatusUnprocessableEntity, "error", errorsMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.service.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register account failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := FormatUser(newUser, "initoken")
	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
