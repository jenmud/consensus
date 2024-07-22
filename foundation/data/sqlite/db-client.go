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

// NewDB creates a new database connection with the schema applied.
func NewDB(dsn string) (*sql.DB, error) {
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

	return db, err
}
