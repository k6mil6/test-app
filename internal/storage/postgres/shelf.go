package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type ShelfStorage struct {
	db *sqlx.DB
}

func NewShelfStorage(db *sqlx.DB) *ShelfStorage {
	return &ShelfStorage{db: db}
}

func (s *ShelfStorage) SelectNameById(ctx context.Context, itemID int) (string, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return "", err
	}

	defer conn.Close()

	var name string

	err = conn.QueryRowxContext(ctx, `SELECT name FROM shelves WHERE id = $1`, itemID).Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}

type dbShelf struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
