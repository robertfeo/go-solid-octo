package database

import (
	"backend/types"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {
	os.Setenv("DB_URL", "postgres://library_db:password@db:5432/library")
	dbUrl := os.Getenv("DB_URL")
	newpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Println("Error while creating db pool")
		return
	}
	Pool = newpool
}

// GetBooks retrieves all books from the database.
func GetBooks() ([]types.Book, error) {
	var books []types.Book
	rows, err := Pool.Query(context.Background(), "SELECT id, title, author, quantity FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b types.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
