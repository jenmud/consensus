// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Event struct {
	ID     int           `json:"id"`
	Op     string        `json:"op"`
	Kind   string        `json:"kind"`
	Fields []*EventField `json:"fields"`
}

type EventField struct {
	OldValue *string `json:"oldValue"`
	NewValue string  `json:"newValue"`
}
