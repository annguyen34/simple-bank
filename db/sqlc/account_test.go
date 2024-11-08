package db

import (
	"context"
	"testing"
	"time"

	"github.com/annguyen34/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Accounts {
	arg := CreateAccountParams{
		Owner:    util.RandomString(10),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account

}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	account_test, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account_test)

	require.Equal(t, account.Owner, account_test.Owner)
	require.Equal(t, account.Balance, account_test.Balance)
	require.Equal(t, account.Currency, account_test.Currency)
	require.Equal(t, account.ID, account_test.ID)
	require.Equal(t, account.CreatedAt, account_test.CreatedAt)

	require.WithinDuration(t, account.CreatedAt, account_test.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	account_test, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account_test)
	require.Equal(t, account.ID, account_test.ID)
	require.Equal(t, arg.Balance, account_test.Balance)
	require.Equal(t, account.Owner, account_test.Owner)
	require.Equal(t, account.Currency, account_test.Currency)
	require.WithinDuration(t, account.CreatedAt, account_test.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account_test, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, account_test)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
