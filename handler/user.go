package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func Userhandler(service user.Service) *userHandler {
	return &userHandler{service}
}

// Register to save data on database
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatUser := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatUser)
	c.JSON(http.StatusOK, response)
}

// Login user to application
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatuser := user.FormatUser(loginUser, "tokentokentoken")
	response := helper.APIResponse("Login successfully", http.StatusOK, "success", formatuser)
	c.JSON(http.StatusOK, response)
}

// Check email availability when user registers
func (h *userHandler) CheckEmailAvaibility(c *gin.Context) {
	var inputEmail user.CheckEmailInput

	if err := c.ShouldBindJSON(&inputEmail); err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": error}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(inputEmail)

	if err != nil {
		errorMessage := gin.H{"error": "Server error"}
		response := helper.APIResponse("Email not available", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
