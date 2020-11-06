package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// Author struct
type Author struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books map[string]Book
var authors map[string]Author

// Get single book
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	books = make(map[string]Book)
	authors = make(map[string]Author)

	authors["1"] = Author{ID: "1", Firstname: "John", Lastname: "Doe"}
	authors["2"] = Author{ID: "2", Firstname: "Steve", Lastname: "Smith"}

	books["1"] = Book{ID: "1", Isbn: "438227", Title: "Book One", Author: authors["1"]}
	books["2"] = Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: authors["2"]}

	// Route handles & endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	// r.HandleFunc("/books", getBooks).Methods("GET")
	// r.HandleFunc("/books/{id}", getBook).Methods("GET")
	// r.HandleFunc("/books", createBook).Methods("POST")
	// r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
