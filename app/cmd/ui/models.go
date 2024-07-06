package ui

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

type Card struct {
	ID      int64
	Title   string
	Content string
}

type Project struct {
	ID          int64
	Title       string
	Description string
	Backlog     []Card
	InProgress  []Card
	CodeReview  []Card
	Testing     []Card
	Done        []Card
	Owner       User
}

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

// JWT represents a JSON Web Token
type JWT struct {
	ID        string    `json:"jti"`  // JWT ID is a unique identifier to prevent replay attacks
	Audience  []string  `json:"aud"`  // Audience is the intended recipient of the JWT
	Subject   int64     `json:"sub"`  // Subject which can be used to identify the user, e.g. user id
	IssuedAt  time.Time `json:"iat"`  // IssuedAt is the time at which the JWT was issued
	ExpiresAt time.Time `json:"exp"`  // ExpiresAt is the time at which the JWT will expire
	User      User      `json:"user"` // User is the user associated with the JWT
}

// AsMap is a helper function to convert a JWT to a map
func (j JWT) AsMap() map[string]any {
	b, err := json.Marshal(j)
	if err != nil {
		return nil
	}

	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil
	}

	return m
}

// JWTFromCtx is a helper function to get the JWT from the context
func JWTFromCtx(ctx context.Context) (JWT, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return JWT{}, err
	}

	return JWTFromMap(claims)
}

// JWTFromMap is a helper function to convert a map to a JWT
func JWTFromMap(m map[string]any) (JWT, error) {
	if m == nil {
		return JWT{}, nil
	}

	b, err := json.Marshal(m)
	if err != nil {
		return JWT{}, err
	}

	var j JWT
	return j, json.Unmarshal(b, &j)
}

// Payload is the main page payload
type Payload struct {
	Project Project
	JWT     JWT
}
