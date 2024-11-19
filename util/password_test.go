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

	err = CheckPasswordHash(password, hashedPassword)
	require.NoError(t, err)

	if err != nil {
		t.Errorf("error checking password hash: %v", err)
	}

	wrongPassword := RandomString(10)
	err = CheckPasswordHash(wrongPassword, hashedPassword)
	require.Error(t, err)
	if err == nil {
		t.Errorf("expect error but got nil")
	}

}
