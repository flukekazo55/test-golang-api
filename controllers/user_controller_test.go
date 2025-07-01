package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	// models import
	"test-golang-api/models"
)

func setupTestRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUserByID)
	r.POST("/users", PostUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)
	return r
}

func TestGetUsers(t *testing.T) {
	router := setupTestRouter()
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestGetUserByID(t *testing.T) {
	router := setupTestRouter()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	router := setupTestRouter()
	req, _ := http.NewRequest("GET", "/users/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 (not found response), got %d", w.Code)
	}
}

func TestPostUser(t *testing.T) {
	router := setupTestRouter()

	newUser := models.User{UserId: "10", UserName: "new_user", Name: "New User"}
	body, _ := json.Marshal(newUser)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestUpdateUser(t *testing.T) {
	router := setupTestRouter()

	updatedUser := models.User{UserName: "updated_user", Name: "Updated Name"}
	body, _ := json.Marshal(updatedUser)

	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestUpdateUser_NotFound(t *testing.T) {
	router := setupTestRouter()

	updatedUser := models.User{UserName: "ghost", Name: "Ghost User"}
	body, _ := json.Marshal(updatedUser)

	req, _ := http.NewRequest("PUT", "/users/999", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	router := setupTestRouter()
	req, _ := http.NewRequest("DELETE", "/users/2", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestDeleteUser_NotFound(t *testing.T) {
	router := setupTestRouter()
	req, _ := http.NewRequest("DELETE", "/users/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}
}
