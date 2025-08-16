package data

import (
	"context"
	"database/sql"
	"time"
)

type Transaction struct {
	TransactionID string    `json:"transactionId"`
	UserID        int64     `json:"-"`
	Amount        float64   `json:"amount"`
	State         string    `json:"state"`
	SourceType    string    `json:"sourceType"`
	CreatedAt     time.Time `json:"createdAt"`
}

type TransactionModel struct {
	DB *sql.DB
}

func (t *TransactionModel) Insert(txn *sql.Tx, transaction *Transaction) error {
	query := `
		INSERT INTO transactions (transaction_id, user_id, amount, state, source_type)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at`

	args := []any{
		transaction.TransactionID,
		transaction.UserID,
		transaction.Amount,
		transaction.State,
		transaction.SourceType,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return txn.QueryRowContext(ctx, query, args...).Scan(&transaction.CreatedAt)
}
