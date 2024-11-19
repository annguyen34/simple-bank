package api

import (
	"testing"

	db "github.com/annguyen34/simple-bank/db/sqlc"
	"github.com/annguyen34/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func randomUser(t *testing.T) (user db.Users, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.Users{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	return
}
