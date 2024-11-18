package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := RandomString(10)
	hashedPassword, err := HashPassword(password)
	require.NotEmpty(t, hashedPassword)
	require.NoError(t, err)

	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}
	if !CheckPasswordHash(password, hashedPassword) {
		t.Errorf("password and hash do not match")
	}

	wrongPassword := RandomString(10)
	if CheckPasswordHash(wrongPassword, hashedPassword) {
		t.Errorf("wrong password match hash")
	}
}
