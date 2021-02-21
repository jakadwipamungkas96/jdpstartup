package handler

import (
	"jdpstartup/helper"
	"jdpstartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//Get input user
	//Map input user -> struct RegisterUserInput
	//Struct kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMsg := gin.H{"errors": errors} // gin.H adalah map, key string, value errors

		response := helper.APIResponse("Register account failed !", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed !", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "testokentokentoken")

	response := helper.APIResponse("Account has been registered !", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMsg := gin.H{"errors": errors} // gin.H adalah map, key string, value errors

		response := helper.APIResponse("Login Failed !", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return // dihentikan
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMsg := gin.H{"errors": err.Error()} // gin.H adalah map, key string, value errors
		response := helper.APIResponse("Login Failed !", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return // dihentikan
	}

	formatter := user.FormatUser(loggedinUser, "testokentokentoken")
	response := helper.APIResponse("Login Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
