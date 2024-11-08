package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			arg := TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			}
			result, err := store.TranferTx(ctx, arg)

			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		require.Equal(t, amount, result.Transfer.Amount)
		require.Equal(t, account1.ID, result.Transfer.FromAccountID)
		require.Equal(t, account2.ID, result.Transfer.ToAccountID)

		// check entries
		_, err = store.GetEntry(context.Background(), result.FromEntry.ID)
		require.NoError(t, err)
		_, err = store.GetEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, err)
	}

}
