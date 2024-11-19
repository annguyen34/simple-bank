package db

import (
	"context"
	"testing"
	"time"

	"github.com/annguyen34/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entries {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testStore.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	entry := createRandomEntry(t)
	require.NotEmpty(t, entry)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)

	entry_test, err := testStore.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry_test)
	require.Equal(t, entry.AccountID, entry_test.AccountID)
	require.Equal(t, entry.Amount, entry_test.Amount)
	require.WithinDuration(t, entry.CreatedAt, entry_test.CreatedAt, time.Second)
}
