package sqlite

import (
	"database/sql"
	"testing"
)

func TestApplySchema(t *testing.T) {
	t.Run("successfully applies the schema", func(t *testing.T) {
		// Create a test database
		db, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		// Apply the schema
		err = ApplySchema(db, "CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("returns an error if the schema cannot be applied", func(t *testing.T) {
		// Create a test database
		db, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		// Apply the schema
		err = ApplySchema(db, "CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT") // missing a closing parenthesis
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestNewClient(t *testing.T) {
	// Test case 1: Successful database connection
	t.Run("Successful database connection", func(t *testing.T) {
		// Call the NewClient function
		client, err := NewClient("file:unittest?mode=memory")

		// Assert the expected output
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		if client == nil {
			t.Error("Expected non-nil Queries struct, got nil")
		}
	})

	// Test case 2: Failed database connection
	t.Run("Failed database connection", func(t *testing.T) {
		// Call the NewClient function
		client, err := NewClient("file:/dev/null")

		// Assert the expected output
		if err == nil {
			t.Error("Expected non-nil error, got nil")
		}

		if client != nil {
			t.Error("Expected nil Queries struct, got non-nil")
		}
	})
}
