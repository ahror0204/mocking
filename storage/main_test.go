package storage

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	strg StorageI
)

func TestMain(m *testing.M) {
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

	strg = NewStorage(db)
	os.Exit(m.Run())
}
