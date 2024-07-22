package crypto

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:    "valid password",
			input:   "password123",
			wantErr: nil,
		},
		{
			name:    "empty password",
			input:   "",
			wantErr: bcrypt.ErrHashTooShort,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, err := HashPassword(tt.input)
			if err != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v %s", err, tt.wantErr, hashedPassword)
				return
			}
			if err == nil && !CheckPasswordHash(hashedPassword, tt.input) {
				t.Errorf("HashPassword() = %v, does not match hashed password", tt.input)
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "password match",
			args: args{
				password: "password123",
				hash:     "$2a$14$bIzeEuctt4uSeZw1o4odeOtEHqN7RTvbX2.B8xL8daSJGvB8VaAym",
			},
			want: true,
		},
		{
			name: "password dont match",
			args: args{
				password: "mismatched-passwords",
				hash:     "$2a$14$bIzeEuctt4uSeZw1o4odeOtEHqN7RTvbX2.B8xL8daSJGvB8VaAym",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordHash(tt.args.hash, tt.args.password); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
