// based off: https://github.com/bradtraversy/go_restapi
package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

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

type RequestError struct {
	Message string `json:"message"`
}

var books []Book
var authors []Author

func getIDFromRequest(r *http.Request) (int, error) {
	params := mux.Vars(r)
	idStr, prs := params["id"]
	if !prs {
		return -1, errors.New("Invalid id")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		return -1, errors.New("Invalid id")
	}
	return id, nil
}

func validateIDRange(id int, lower int, upper int) error {
	if id < lower || id > upper {
		return errors.New("Id out of range")
	}
	return nil
}

// Get single book
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

	reqError := RequestError{Message: "Unable to get the book"}
	w.Header().Set("Content-Type", "application/json")

	id, err := getIDFromRequest(r)
	if err != nil {
		log.Println("Invalid id")
		json.NewEncoder(w).Encode(reqError)
		return
	}
	rangeErr := validateIDRange(id, 0, len(books))
	if rangeErr != nil {
		log.Println("Invalid id")
		json.NewEncoder(w).Encode(reqError)
		return
	}
	book := books[id]
	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	reqError := RequestError{Message: "Unable to create book"}
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		json.NewEncoder(w).Encode(reqError)
		return
	}
	book.ID = strconv.Itoa(len(books))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	authors = append(authors, Author{ID: "0", Firstname: "John", Lastname: "Doe"})
	authors = append(authors, Author{ID: "1", Firstname: "Steve", Lastname: "Smith"})

	books = append(books, Book{ID: "0", Isbn: "438227", Title: "Book One", Author: authors[0]})
	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: authors[1]})

	// Route handles & endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	// r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
