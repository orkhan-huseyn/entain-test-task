package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type User struct {
	Id        uint64    `json:"userId"`
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

	err := u.DB.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Balance, &user.CreatedAt)

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
