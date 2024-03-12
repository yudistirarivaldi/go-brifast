package handler

import (
	"brifast-service-login/auth"
	"brifast-service-login/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService auth.Service
}

func NewAuthHandler(authService auth.Service) *authHandler {
	return &authHandler{authService}
}

func (h *authHandler) Login(c *gin.Context) {

	var input auth.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{ "errors" : errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.authService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()} 
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response) 
		return
	}

	formatter := auth.FormatUser(loggedinUser)

	response := helper.APIResponse("Login Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}