package data

import "database/sql"

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Get(id uint64) (any, error) {
	return nil, nil
}
