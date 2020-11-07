// based off: https://github.com/bradtraversy/go_restapi
package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Book struct (Model)
type Book struct {
	ID     int    `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// func idValidationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		_, err := getIDFromRequest(r)
// 		if err == nil {
// 			next.ServeHTTP(w, r)
// 		} else {
// 			log.Println("Invalid id")
// 			http.Error(w, "Invalid id", http.StatusBadRequest)
// 		}

// 		// Call the next handler, which can be another middleware in the chain, or the final handler.
// 	})
// }

// Get single book
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := `SELECT id, Isbn, Title, Author FROM Book`
	rows, err := db.Query(q)
	if err != nil {
		http.Error(w, "Cannot get books", http.StatusInternalServerError)
	}
	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		var id int
		var isbn, title, author string
		if err := rows.Scan(&id, &isbn, &title, &author); err != nil {
			log.Fatal(err)
		}
		books = append(books, Book{ID: id, Isbn: isbn, Title: title, Author: author})

	}

	json.NewEncoder(w).Encode(books)
}

// func getBook(w http.ResponseWriter, r *http.Request) {

// 	reqError := RequestError{Message: "Unable to get the book"}
// 	w.Header().Set("Content-Type", "application/json")

// 	id, err := getIDFromRequest(r)
// 	if err != nil {
// 		log.Println("Invalid id")
// 		json.NewEncoder(w).Encode(reqError)
// 		return
// 	}
// 	rangeErr := validateIDRange(id, 0, len(books))
// 	if rangeErr != nil {
// 		log.Println("Invalid id")
// 		json.NewEncoder(w).Encode(reqError)
// 		return
// 	}
// 	book := books[id]
// 	json.NewEncoder(w).Encode(book)
// }

// func createBook(w http.ResponseWriter, r *http.Request) {
// 	reqError := RequestError{Message: "Unable to create book"}
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	err := json.NewDecoder(r.Body).Decode(&book)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(reqError)
// 		return
// 	}
// 	book.ID = strconv.Itoa(len(books))
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)
// }

// TODO: use dependency injection
var db *sql.DB

func use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	// authors = append(authors, Author{ID: "0", Firstname: "John", Lastname: "Doe"})
	// authors = append(authors, Author{ID: "1", Firstname: "Steve", Lastname: "Smith"})

	// books = append(books, Book{ID: "0", Isbn: "438227", Title: "Book One", Author: authors[0]})
	// books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: authors[1]})

	db, err := sql.Open("sqlite3", "./books.db")
	use(db)

	if err != nil {
		log.Fatal(err)
	}

	// Route handles & endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	// r.HandleFunc("/books/{id}", getBook).Methods("GET")
	// r.HandleFunc("/books", createBook).Methods("POST")
	// r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	r.Use(loggingMiddleware)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
