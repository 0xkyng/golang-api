package main

import (
	"errors"
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

//////////////////////////////////////////////
// Get book by id
func getBookById(id string) (*book, error) {
	for i, b := range books{
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func bookById(context *gin.Context) {
	id := context.Param("id")
	book, err := getBookById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, book)

}
////////////////////////////////////////////

/////////////////////////////////////////
// Checkout book
func checkOutBook(context *gin.Context) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query"})
	}

	book, err := getBookById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"messsge": "Book not found"})
	return
	}
	

	if book.Quantity <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not avqilable"})
		return
	}

	book.Quantity -= 1
	context.IndentedJSON(http.StatusOK, book)
}
/////////////////////////////////////////////////

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func returnBook(context *gin.Context) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query"})
	}

	book, err := getBookById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"messsge": "Book not found"})
	return
	}

	book.Quantity += 1
	context.IndentedJSON(http.StatusOK, book)

}
//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

////////////////////////////////////
// POST Method
func addBook(context *gin.Context) {
	var newBook book

	err := context.BindJSON(&newBook)
	if err != nil{
		return
	}

	books = append(books, newBook)

	context.IndentedJSON(http.StatusCreated, newBook)
}
////////////////////////////////////


func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", addBook)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}