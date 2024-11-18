package db

import (
	"context"
	"testing"
	"time"

	"github.com/annguyen34/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) Users {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.CreatedAt)

	return user

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	user_test, err := testStore.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user_test)

	require.Equal(t, user.Username, user_test.Username)
	require.Equal(t, user.FullName, user_test.FullName)
	require.Equal(t, user.HashedPassword, user_test.HashedPassword)
	require.Equal(t, user.Email, user_test.Email)
	require.WithinDuration(t, user.CreatedAt, user_test.CreatedAt, time.Second)

}
