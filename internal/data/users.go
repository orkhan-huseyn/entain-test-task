package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID        uint64
	Balance   float64
	CreatedAt time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Get(id uint64) (*User, error) {
	var user User
	query := `
		SELECT id, balance, created_at
		FROM users
		WHERE id=$1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Balance, &user.CreatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserModel) GetForUpdate(txn *sql.Tx, id uint64) (*User, error) {
	var user User
	query := `
		SELECT id, balance, created_at
		FROM users
		WHERE id=$1
		FOR UPDATE`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := txn.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Balance, &user.CreatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserModel) Update(txn *sql.Tx, id uint64, balanceIncrement float64) error {
	query := `
		UPDATE users
		SET balance = balance + $1
		WHERE id = $2`

	args := []any{balanceIncrement, id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := txn.QueryRowContext(ctx, query, args...).Err()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		default:
			return err
		}
	}

	return nil
}
