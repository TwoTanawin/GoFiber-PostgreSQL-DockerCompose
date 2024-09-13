// db_test.go
package main

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestDatabaseConnection(t *testing.T) {
	connStr := "user=admin password=password dbname=movie_app host=localhost port=5433 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	t.Logf("Successfully connected to the database")
}

