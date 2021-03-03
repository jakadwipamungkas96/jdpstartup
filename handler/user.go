package handler

import (
	"fmt"
	"jdpstartup/auth"
	"jdpstartup/helper"
	"jdpstartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
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

	token, err := h.authService.GenerateToken(newUser.Id)
	if err != nil {
		response := helper.APIResponse("Register account failed !", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

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
	token, err := h.authService.GenerateToken(loggedinUser.Id)
	if err != nil {
		response := helper.APIResponse("Login failed !", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Login Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMsg := gin.H{"errors": errors} // gin.H adalah map, key string, value errors

		response := helper.APIResponse("Email checking failed !", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return // dihentikan
	}

	checkEmailAvailable, err := h.userService.CheckAvailableEmail(input)
	if err != nil {
		errorMsg := gin.H{"errors": "Server Error"} // gin.H adalah map, key string, value errors
		response := helper.APIResponse("Email checking failed !", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return // dihentikan
	}

	data := gin.H{
		"is_ava": checkEmailAvailable,
	}

	metaMsg := "Email has been registered"
	if checkEmailAvailable {
		metaMsg = "Email is available"
	}

	response := helper.APIResponse(metaMsg, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload ava image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get from JWT
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.Id

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload ava image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAva(userID, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload ava image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Ava succesfull uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
