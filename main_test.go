package main

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"

	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	return router
}
func TestAddTodo(t *testing.T) {
	r := SetUpRouter()
	r.POST("/todos", addTodo)

	todoRequest := todo{
		ID:        "100",
		Item:      "Goto Sleep",
		Completed: false,
	}
	jsonValue, _ := json.Marshal(todoRequest)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var todoResponse todo
	json.Unmarshal(w.Body.Bytes(), &todoResponse)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, todoResponse)
}

func TestGetTodos(t *testing.T) {
	r := SetUpRouter()
	r.GET("/todos", getTodos)
	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var todos []todo
	json.Unmarshal(w.Body.Bytes(), &todos)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, todos)
}

func TestGetTodo(t *testing.T) {
	r := SetUpRouter()
	r.GET("/todos/:id", getTodo)
	req, _ := http.NewRequest("GET", "/todos/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var todo todo
	json.Unmarshal(w.Body.Bytes(), &todo)
	assert.Equal(t, http.StatusOK, w.Code)

	assert.NotEmpty(t, todo)
}

func TestGetTodoNotFound(t *testing.T) {
	r := SetUpRouter()
	r.GET("/todos/:id", getTodo)
	req, _ := http.NewRequest("GET", "/todos/5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var todo todo
	json.Unmarshal(w.Body.Bytes(), &todo)
	assert.Equal(t, http.StatusNotFound, w.Code)

	assert.Empty(t, todo)
}

func TestToggleTodoStatus(t *testing.T) {
	r := SetUpRouter()
	r.GET("/todos/:id", toggleTodoStatus)
	req, _ := http.NewRequest("GET", "/todos/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var todo todo
	json.Unmarshal(w.Body.Bytes(), &todo)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, todo)
}

func TestToggleTodoStatusNotFound(t *testing.T) {
	r := SetUpRouter()
	r.GET("/todos/:id", toggleTodoStatus)
	req, _ := http.NewRequest("GET", "/todos/10", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var todo todo
	json.Unmarshal(w.Body.Bytes(), &todo)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Empty(t, todo)
}
