package controllers

import (
	"gin-flemarket/dto"
	"gin-flemarket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	SingUp(ctx *gin.Context)
}

type AuthController struct {
	services services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{services: service}
}

func (c *AuthController) SingUp(ctx *gin.Context) {
	var input dto.SignUpInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := c.services.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to create user"})
		return
	}

	ctx.Status(http.StatusCreated)
}
