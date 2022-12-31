package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// book struct for our library
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// Slice / list of books
var books = []book{
	{ID: "1", Title: "Landing your first remot job", Author: "Isaac", Quantity: 5},
	{ID: "2", Title: "Working with FANG", Author: "codekyng", Quantity: 4},
	{ID: "3", Title: "Transitioning to tech", Author: "Izik", Quantity: 10},
}

/////////////////////////////////////
// GET Method
// getBooks handles the route of getting
// All the different books

func getBooks(context *gin.Context) {
	// Gin.context contains all the information about the incoming request
   // It allows you to return a response
	context.IndentedJSON(http.StatusOK, books) // Transform the data from the request
  // And return it as nicely formatted json that's properly indented
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
// End point for adding a book to list of books 
func addBook(context *gin.Context) {
	// Firt create a new book
	// Get the data of the new book to be added from the request
	var newBook book

	// Bind the json from the request body
	// To the newBook with book type
	err := context.BindJSON(&newBook)
	if err != nil{
		return
	}

	// Add the newBook to the books list
	books = append(books, newBook)

	// Return the newBook
	context.IndentedJSON(http.StatusCreated, newBook)
}
////////////////////////////////////


func main() {
	// Set up router
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", addBook)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}