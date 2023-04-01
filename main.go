package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, todos)
}

func addTodo(ctx *gin.Context) {
	var newTodo todo
	if err := ctx.BindJSON(&newTodo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	todos = append(todos, newTodo)
	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {

	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func toggleTodoStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	ctx.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)

	router.Run("localhost:9090")
}
