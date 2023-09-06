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

func (h *userHandler) Login(c *gin.Context) {
	var input LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationsErrors(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogin, err := h.service.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := FormatUser(userLogin, "initoken")

	response := helper.ApiResponse("Successfully login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
