/*
based on: https://github.com/bradtraversy/go_restapi
TODO:
- Write tests
- Use dependency injection for logging, db etc
- Clean up error handling
- Look at ORMs
- Research input validation
- DB indices/constraints etc.
*/
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID     int64  `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func convertBookRows(rows *sql.Rows) []Book {

	books := make([]Book, 0)
	for rows.Next() {
		var id int64
		var isbn, title, author string
		if err := rows.Scan(&id, &isbn, &title, &author); err != nil {
			log.Fatal(err)
		}
		books = append(books, Book{ID: id, Isbn: isbn, Title: title, Author: author})

	}
	return books
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := `SELECT id, Isbn, Title, Author FROM Book`
	rows, err := db.Query(q)
	defer rows.Close()
	if err != nil {
		http.Error(w, "Cannot get books", http.StatusInternalServerError)
		return
	}
	books := convertBookRows(rows)

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	q := `SELECT id, Isbn, Title, Author FROM Book WHERE id=?`
	rows, err := db.Query(q, id)
	defer rows.Close()

	if err != nil {
		http.Error(w, "Cannot get books", http.StatusInternalServerError)
		return
	}

	books := convertBookRows(rows)
	rows.Close()
	if len(books) == 1 {
		json.NewEncoder(w).Encode(books[0])
	} else {
		http.Error(w, "Cannot find book.", http.StatusNotFound)
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println(err)
		http.Error(w, "Cannot update book.", http.StatusBadRequest)
		return
	}

	q := `UPDATE Book SET Isbn=?, Title=?, Author=? WHERE id=?`
	res, err := db.Exec(q, book.Isbn, book.Title, book.Author, id)

	if err != nil {
		http.Error(w, "Cannot update book", http.StatusInternalServerError)
		return
	}

	book.ID, _ = res.LastInsertId()
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	q := `DELETE FROM Book WHERE id=?`
	res, _ := db.Exec(q, id)

	if rowsDeleted, _ := res.RowsAffected(); rowsDeleted != 1 {
		http.Error(w, "Cannot delete book", http.StatusInternalServerError)
		return
	}

	log.Println(res)
	w.WriteHeader(http.StatusOK)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println(err)
		http.Error(w, "Cannot create book.", http.StatusBadRequest)
		return
	}
	q := `INSERT INTO Book (Isbn, Title, Author) VALUES (?, ?, ?)`
	res, err := db.Exec(q, book.Isbn, book.Title, book.Author)
	if err != nil {
		log.Println(err)
		http.Error(w, "Cannot create book.", http.StatusBadRequest)
		return
	}
	book.ID, _ = res.LastInsertId()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)

}

var db *sql.DB

func main() {
	r := mux.NewRouter()

	var err error
	db, err = sql.Open("sqlite3", "./books.db")

	if err != nil {
		log.Fatal(err)
	}

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	r.Use(loggingMiddleware)

	log.Fatal(http.ListenAndServe(":8000", r))
}
