package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB // db is the database connection pool
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,      // => for transaction
		Queries: New(db), // => for queries
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil) // => nil is the default isolation level
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferResult struct {
	Transfer    Transfers `json:"transfer"`
	FromAccount Accounts  `json:"from_account"`
	ToAccount   Accounts  `json:"to_account"`
	FromEntry   Entries   `json:"from_entry"`
	ToEntry     Entries   `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other
// It creates a transfer record => add account entries => update account balance
func (store *Store) TranferTx(ctx context.Context, arg TransferTxParams) (TransferResult, error) {
	var result TransferResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Get the account that will send money
		result.FromAccount, err = q.GetAccount(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}

		// Get the account that will receive money
		result.ToAccount, err = q.GetAccount(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}

		// Create a transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create an entry for the send accout
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create an entry for the receive account
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return TransferResult{}, err
	}

	return result, nil
}
