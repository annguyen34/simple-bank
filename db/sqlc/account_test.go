package db

import (
	"context"
	"testing"
	"time"

	"github.com/annguyen34/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Accounts {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testStore.CreateAccount(context.Background(), arg)
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

	account_test, err := testStore.GetAccount(context.Background(), account.ID)
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

	account_test, err := testStore.UpdateAccount(context.Background(), arg)
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

	err := testStore.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account_test, err := testStore.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, account_test)
}

func TestListAccount(t *testing.T) {
	var lastAccount Accounts
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}

}
