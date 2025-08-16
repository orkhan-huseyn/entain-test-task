package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID        uint64    `json:"userId"`
	Balance   string    `json:"balance"`
	CreatedAt time.Time `json:"-"`
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Get(id uint64) (*User, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

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

func (u *UserModel) Update(id uint64, balanceIncrement float64) error {
	query := `
		UPDATE users
		SET balance = balance + $1
		WHERE id = $2`

	args := []any{balanceIncrement, id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.DB.QueryRowContext(ctx, query, args...).Err()
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
