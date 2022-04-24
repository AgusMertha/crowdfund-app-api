package handler

import (
	"fmt"
	"kitabantu-api/auth"
	"kitabantu-api/helper"
	"kitabantu-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.UserService
	authService auth.Service
}

func NewUserHandler(userService user.UserService, authService auth.Service) *UserHandler {
	return &UserHandler{userService, authService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	// get user input
	// mapping input to struct UserInput
	// pass as service param

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}
		errorResponse := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		errorResponse := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// generate token
	token, err := h.authService.GenerateToken(newUser.Id)

	if err != nil {
		errorResponse := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	formatter := user.FormatUser(newUser, token)
	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	// get email & password
	var input user.LoginUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}
		errorResponse := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	userLogin, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errorResponse := helper.ApiResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// generate token
	token, err := h.authService.GenerateToken(userLogin.Id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errorResponse := helper.ApiResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	formatter := user.FormatUser(userLogin, token)
	response := helper.ApiResponse("Login success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}
		errorResponse := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {

		errorMessage := gin.H{"errors": "Server error"}
		errorResponse := helper.ApiResponse("Email checking failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	message := "Email has been registered"
	if isEmailAvailable {
		message = "Email is available"
	}
	response := helper.ApiResponse(message, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {

		errorMessage := gin.H{"is_uploaded": false}
		errorResponse := helper.ApiResponse("Failed to upload avatar image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	// save avatar string to DB
	userId := 9
	path := fmt.Sprintf("images/profile/%d-%s", userId, file.Filename)

	// save file to directory
	err = c.SaveUploadedFile(file, path)

	if err != nil {

		errorMessage := gin.H{"is_uploaded": false}
		errorResponse := helper.ApiResponse("Failed to upload avatar image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	_, err = h.userService.SaveAvatar(userId, path)

	if err != nil {

		errorMessage := gin.H{"is_uploaded": false}
		errorResponse := helper.ApiResponse("Failed to upload avatar image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errorResponse := helper.ApiResponse("Success to upload avatar image", http.StatusOK, "success", gin.H{"is_uploaded": true})
	c.JSON(http.StatusOK, errorResponse)
}
