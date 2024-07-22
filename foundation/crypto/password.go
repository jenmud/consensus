package crypto

import (
	rand "math/rand/v2"

	"golang.org/x/crypto/bcrypt"
)

// Secret is a generated secret key used for OAuth2.
func Secret() string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789*%&^$#@!~")
	b := make([]rune, 32)
	for i := range b {
		b[i] = chars[rand.IntN(len(chars))]
	}
	return string(b)
}

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", bcrypt.ErrHashTooShort
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a bcrypt hash to a password
func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
