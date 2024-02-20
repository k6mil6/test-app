package postgres

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/k6mil6/test-app/internal/model"
)

type OrderStorage struct {
	db *sqlx.DB
}

func NewOrderStorage(db *sqlx.DB) *OrderStorage {
	return &OrderStorage{db: db}
}

func (o *OrderStorage) Create(ctx context.Context, order model.Order) error {
	tx, err := o.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	orderResult, err := tx.ExecContext(ctx, `INSERT INTO orders (user_id) VALUES ($1) RETURNING id`, order.UserID)
	if err != nil {
		return err
	}

	orderID, err := orderResult.LastInsertId()
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		_, err := tx.ExecContext(ctx, `INSERT INTO orders_items (order_id, item_id) VALUES ($1, $2)`, orderID, item.ID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderStorage) Select(ctx context.Context, orderID int) (model.Order, error) {
	tx, err := o.db.BeginTxx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return model.Order{}, err
	}
	defer tx.Rollback()

	var dbOrder dbOrder
	err = tx.GetContext(ctx, &dbOrder, "SELECT id, user_id FROM orders WHERE id = $1", orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Order{}, nil
		}
		return model.Order{}, err
	}

	var dbItems []dbItem
	err = tx.SelectContext(ctx, &dbItems, "SELECT i.id, i.name, i.main_shelf_id FROM items i INNER JOIN orders_items oi ON i.id = oi.item_id WHERE oi.order_id = $1", dbOrder.ID)
	if err != nil {
		return model.Order{}, err
	}

	var order model.Order
	order.ID = dbOrder.ID
	order.UserID = dbOrder.UserID

	for _, dbItem := range dbItems {
		var item model.Item
		item.ID = dbItem.ID
		item.Name = dbItem.Name
		item.MainShelfID = dbItem.MainShelfID

		var dbShelves []dbShelf
		err = tx.SelectContext(ctx, &dbShelves, `SELECT s.id, s.name FROM shelves s INNER JOIN items_shelves "is" ON s.id = "is".shelf_id WHERE "is".item_id = $1`, dbItem.ID)
		if err != nil {
			return model.Order{}, err
		}

		for _, dbShelf := range dbShelves {
			var shelf model.Shelf
			shelf.ID = dbShelf.ID
			shelf.Name = dbShelf.Name
			item.Shelves = append(item.Shelves, shelf)
		}

		order.Items = append(order.Items, item)
	}

	err = tx.Commit()
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

type dbOrder struct {
	ID     int `db:"id"`
	UserID int `db:"user_id"`
	Items  []dbItem
}
