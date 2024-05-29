package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// ApplySchema applies the schema to the given database.
//
// It takes a *sql.DB as a parameter, which represents the database connection.
// The function reads the schema from the "schema.sql" file using the Schema embed.FS.
// It then begins a transaction on the database, executes the schema SQL statements,
// and commits the transaction. If any error occurs during the process, it is returned.
//
// Returns an error if there was a problem reading the schema file or executing the SQL statements.
func ApplySchema(db *sql.DB, schema string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(schema)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// NewClient creates a new client for interacting with a SQLite database.
//
// It takes a DSN (Data Source Name) string as a parameter, which specifies the
// connection details for the database. The DSN should be in the format:
// `<filename>|<connection_string>`.
//
// The function returns a pointer to a Queries struct and an error. The Queries
// struct provides methods for executing SQL queries against the database. The
// error is non-nil if there was an error opening the database connection or
// applying the schema.
func NewClient(dsn string) (*Queries, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	content, err := schema.ReadFile("schema.sql")
	if err != nil {
		return nil, err
	}

	if err := ApplySchema(db, string(content)); err != nil {
		return nil, err
	}

	return New(db), nil
}
