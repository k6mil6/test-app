package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type ItemStorage struct {
	db *sqlx.DB
}

func NewItemStorage(db *sqlx.DB) *ItemStorage {
	return &ItemStorage{db: db}
}

func (i *ItemStorage) SelectNameById(ctx context.Context, itemID int) (string, error) {
	conn, err := i.db.Connx(ctx)
	if err != nil {
		return "", err
	}

	defer conn.Close()

	var name string

	err = conn.QueryRowxContext(ctx, `SELECT name FROM items WHERE id = $1`, itemID).Scan(&name)

	return "", err
}

type dbItem struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	MainShelfID int    `db:"main_shelf_id"`
	Quantity    int    `db:"quantity"`
	Shelves     []dbShelf
}
