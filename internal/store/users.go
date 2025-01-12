package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	db *sql.DB
}

func (us *UsersStore) Create(ctx context.Context) error {
	return nil
}
