package ui

import (
	"errors"
	"net/http"

	"github.com/go-chi/oauth"
	"github.com/jenmud/consensus/business/service"
)

// UserAuthVerifier provides user credentials verifier for testing.
type UserAuthVerifier struct {
	client service.ConsensusClient
}

// NewUserAuthVerifier creates a new UserAuthVerifier.
func NewUserAuthVerifier(client service.ConsensusClient) *UserAuthVerifier {
	return &UserAuthVerifier{
		client: client,
	}
}

// ValidateUser validates username and password returning an error if the user credentials are wrong
func (u *UserAuthVerifier) ValidateUser(username, password, scope string, r *http.Request) error {
	_, err := u.client.AuthenticateUser(r.Context(), &service.AuthReq{
		Email:    username,
		Password: password,
	})

	return err
}

// ValidateClient validates clientID and secret returning an error if the client credentials are wrong
func (u *UserAuthVerifier) ValidateClient(clientID, clientSecret, scope string, r *http.Request) error {
	return errors.New("not implemented")
}

// ValidateCode validates token ID
func (u *UserAuthVerifier) ValidateCode(clientID, clientSecret, code, redirectURI string, r *http.Request) (string, error) {
	return "", errors.New("not implemented")
}

// AddClaims provides additional claims to the token
func (u *UserAuthVerifier) AddClaims(tokenType oauth.TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {
	claims := make(map[string]string)
	return claims, nil
}

// AddProperties provides additional information to the token response
func (u *UserAuthVerifier) AddProperties(tokenType oauth.TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {
	props := make(map[string]string)
	return props, nil
}

// ValidateTokenID validates token ID
func (u *UserAuthVerifier) ValidateTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	return nil
}

// StoreTokenID saves the token id generated for the user
func (u *UserAuthVerifier) StoreTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	return nil
}
