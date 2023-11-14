package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {
	dbUrl := "postgres://library_db:password@127.0.0.1:5432/library"
	newpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Println("Error while creating db pool")
		return
	}
	Pool = newpool
}
