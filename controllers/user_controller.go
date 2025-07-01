package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// utils import
	utils "test-golang-api/utils"

	// models import
	models "test-golang-api/models"
)

var users = []models.User{
	{UserId: "1", UserName: "john_doe", Name: "John Doe"},
	{UserId: "2", UserName: "jane_doe", Name: "Jane Doe"},
	{UserId: "3", UserName: "sam_smith", Name: "Sam Smith"},
}

// GET /users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ApiResponse{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    users,
	})
}

// GET /users/:id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.UserId == id {
			c.JSON(http.StatusOK, utils.ApiResponse{
				Status:  http.StatusOK,
				Message: http.StatusText(http.StatusOK),
				Data:    user,
			})

			return
		}
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Status:  http.StatusNotFound,
		Message: "user not found",
		Data:    nil,
	})
}

// POST /users
func PostUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})

		return
	}

	users = append(users, newUser)

	c.JSON(http.StatusCreated, utils.ApiResponse{
		Status:  http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    newUser,
	})
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User

	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})

		return
	}

	for i, u := range users {
		if u.UserId == id {
			users[i].UserName = updatedUser.UserName
			users[i].Name = updatedUser.Name

			c.JSON(http.StatusOK, utils.ApiResponse{
				Status:  http.StatusOK,
				Message: http.StatusText(http.StatusOK),
				Data:    users[i],
			})

			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})

	c.JSON(http.StatusNotFound, utils.ApiResponse{
		Status:  http.StatusNotFound,
		Message: "user not found",
		Data:    nil,
	})
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range users {
		if u.UserId == id {
			deletedUser := u
			users = append(users[:i], users[i+1:]...)

			c.JSON(http.StatusOK, utils.ApiResponse{
				Status:  http.StatusOK,
				Message: http.StatusText(http.StatusOK),
				Data:    deletedUser,
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, utils.ApiResponse{
		Status:  http.StatusNotFound,
		Message: "user not found",
		Data:    nil,
	})
}
