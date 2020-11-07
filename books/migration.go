package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database, _ := sql.Open("sqlite3", "./books.db")
	tx, err := database.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	statement, _ := tx.Prepare(`DROP TABLE Book`)
	statement.Exec()
	statement, _ = tx.Prepare(`
		CREATE TABLE Book (
			ID INTEGER PRIMARY KEY, 
			Isbn TEXT NOT NULL, 
			Title TEXT NOT NULL,
			Author TEXT)
		`)
	statement.Exec()
	statement, _ = tx.Prepare("INSERT INTO Book (Isbn, Title, Author) VALUES (?, ?, ?)")

	books := []struct {
		Isbn   string
		Title  string
		Author string
	}{
		{"234234", "Go Microservices", "Tanzim Mokammel"},
		{"234235", "Go CLI", "Steve Smith"},
		{"234236", "Go Systems", "John Doe"},
	}

	for _, book := range books {
		if _, err := statement.Exec(book.Isbn, book.Title, book.Title); err != nil {
			log.Fatal(err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
