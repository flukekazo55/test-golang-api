package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// utils import
	utils "test-golang-api/utils"
)

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid login data",
			Data:    nil,
		})

		return
	}

	// mock user login
	if req.UserName != "admin" || req.Password != "1234" {
		c.JSON(http.StatusUnauthorized, utils.ApiResponse{
			Status:  http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
			Data:    nil,
		})

		return
	}

	// generate token
	token, err := utils.GenerateToken(req.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Could not create token",
			Data:    nil,
		})

		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Status:  http.StatusOK,
		Message: "Login success",
		Data:    token,
	})
}
