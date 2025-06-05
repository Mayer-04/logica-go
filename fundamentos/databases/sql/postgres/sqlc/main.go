package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"

	"databases/sql/postgres/sqlc/db"
)

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres://mayer:123456@localhost?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	db := db.New(conn)

	author, err := db.GetAuthor(context.Background(), 1)
	if err != nil {
		log.Fatalf("GetAuthor failed: %v\n", err)
	}

	fmt.Println("author:", author.Name)
}
