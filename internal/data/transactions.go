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

func (t *TransactionModel) Insert(txn *Transaction) error {
	query := `
		INSERT INTO transactions (transaction_id, user_id, amount, state, source_type)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at`

	args := []any{txn.TransactionID, txn.UserID, txn.Amount, txn.State, txn.SourceType}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return t.DB.QueryRowContext(ctx, query, args...).Scan(&txn.CreatedAt)
}
