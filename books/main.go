// based off: https://github.com/bradtraversy/go_restapi
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Book struct
type Book struct {
	ID     int    `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func convertBookRows(rows *sql.Rows) []Book {

	books := make([]Book, 0)
	for rows.Next() {
		var id int
		var isbn, title, author string
		if err := rows.Scan(&id, &isbn, &title, &author); err != nil {
			log.Fatal(err)
		}
		books = append(books, Book{ID: id, Isbn: isbn, Title: title, Author: author})

	}
	return books
}

// Get single book
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := `SELECT id, Isbn, Title, Author FROM Book`
	rows, err := db.Query(q)
	if err != nil {
		http.Error(w, "Cannot get books", http.StatusInternalServerError)
	}
	books := convertBookRows(rows)
	rows.Close()

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, prs := params["id"]

	if !prs {
		http.Error(w, "Please provide id", http.StatusBadRequest)

	}
	q := `SELECT id, Isbn, Title, Author FROM Book WHERE id=?`
	rows, err := db.Query(q, id)
	if err != nil {
		http.Error(w, "Cannot get books", http.StatusInternalServerError)
	}

	books := convertBookRows(rows)
	rows.Close()
	if len(books) == 1 {
		json.NewEncoder(w).Encode(books[0])
	} else {
		http.Error(w, "Cannot find book.", http.StatusNotFound)
	}
}

// TODO: use dependency injection
var db *sql.DB

// Main function
func main() {
	r := mux.NewRouter()

	var err error
	db, err = sql.Open("sqlite3", "./books.db")

	if err != nil {
		log.Fatal(err)
	}

	// Route handles & endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	// r.HandleFunc("/books", createBook).Methods("POST")
	// r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	r.Use(loggingMiddleware)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
