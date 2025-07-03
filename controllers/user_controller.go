package controllers

import (
	"context"
	"net/http"
	"time"

	// import gin
	"github.com/gin-gonic/gin"

	// import mongo
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	// import database
	"test-golang-api/database"
	// import models
	"test-golang-api/models"
	// import utils
	"test-golang-api/utils"
)

//

// GET /users
func GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := database.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{Status: 500, Message: "DB error"})
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{Status: 500, Message: "Parse error"})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Status: 200, Message: "OK", Data: users})
}

// GET /users/:id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := database.UserCollection.FindOne(ctx, bson.M{"userId": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, utils.ApiResponse{Status: 404, Message: "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, utils.ApiResponse{Status: 500, Message: "DB error"})
		}
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Status: 200, Message: "OK", Data: user})
}

// POST /users
func PostUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{Status: 400, Message: "Invalid input", Data: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := database.UserCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{Status: 500, Message: "Insert failed"})
		return
	}

	c.JSON(http.StatusCreated, utils.ApiResponse{Status: 201, Message: "User created", Data: newUser})
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updated models.User

	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{Status: 400, Message: "Invalid input", Data: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": id}
	update := bson.M{"$set": bson.M{
		"userName": updated.UserName,
		"name":     updated.Name,
	}}

	res, err := database.UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{Status: 500, Message: "Update failed"})
		return
	}

	if res.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, utils.ApiResponse{Status: 404, Message: "user not found"})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Status: 200, Message: "User updated"})
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := database.UserCollection.DeleteOne(ctx, bson.M{"userId": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{Status: 500, Message: "Delete failed"})
		return
	}

	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, utils.ApiResponse{Status: 404, Message: "user not found"})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Status: 200, Message: "User deleted"})
}
