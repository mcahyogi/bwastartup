package user

import (
	"bwastartup/helper"
	"fmt"
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

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationsErrors(err)
		errMsg := gin.H{"errors": errors}
		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	isEmailAvail, err := h.service.IsEmailAvailable(input)
	if err != nil {
		errMsg := gin.H{"errors": "Server error"}
		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{"is_avalailable": isEmailAvail}
	metaMsg := "Email has been registered"
	if isEmailAvail {
		metaMsg = "Email is available"
	}

	response := helper.ApiResponse(metaMsg, http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId := "1"
	path := fmt.Sprintf("images/%s-%s", userId, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Successfully uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
