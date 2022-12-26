package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Landing your first remot job", Author: "Isaac", Quantity: 5},
	{ID: "2", Title: "Working with FANG", Author: "codekyng", Quantity: 4},
	{ID: "3", Title: "Transitioning to tech", Author: "Izik", Quantity: 10},
}

/////////////////////////////////////
// GET Method
func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}
//////////////////////////////////////


func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}