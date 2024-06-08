// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlite

import (
	"database/sql"
	"time"
)

type Epic struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description sql.NullString
	UserID      int64
}

type Project struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description sql.NullString
	UserID      int64
}

type User struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	FirstName string
	LastName  string
	Password  string
	Role      string
}
