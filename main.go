package main

import (
	"fmt"
	"log"

	"github.com/ahror0204/mocking/storage"
	"github.com/ahror0204/mocking/api"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"postgres",
		"mocking",
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	strg := storage.NewStorage(db)

	server := api.NewServer(strg)

	err = server.Router.Run(":8000")
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
