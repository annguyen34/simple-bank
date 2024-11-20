package token

import (
	"testing"
	"time"

	"github.com/annguyen34/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	if err != nil {
		t.Error(err)
	}

	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expriedAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expriedAt, payload.ExpiredAt, time.Second)
	require.NotZero(t, payload.ID)
}
