package initialisation

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func Init(ctx context.Context, db *sqlx.DB) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	query := "INSERT INTO users (name) VALUES ('Kamil');"
	if _, err := conn.ExecContext(ctx, query); err != nil {
		return err
	}

	query1 := `INSERT INTO shelves (name) VALUES ('А'), ('Б'), ('В'), ('Ж'), ('З');`
	if _, err := conn.ExecContext(ctx, query1); err != nil {
		return err
	}

	query2 := `INSERT INTO items (id, name, main_shelf_id) VALUES
(1, 'Ноутбук', (SELECT id FROM shelves WHERE name = 'А')),
(2, 'Телевизор', (SELECT id FROM shelves WHERE name = 'А')),
(3, 'Телефон', (SELECT id FROM shelves WHERE name = 'Б')),
(4, 'Системный блок', (SELECT id FROM shelves WHERE name = 'Ж')),
(5, 'Часы', (SELECT id FROM shelves WHERE name = 'Ж')),
(6, 'Микрофон', (SELECT id FROM shelves WHERE name = 'Ж'));`

	if _, err := conn.ExecContext(ctx, query2); err != nil {
		return err
	}

	query3 := `INSERT INTO orders (id, user_id) VALUES
(10, 1),
(11, 1),
(14, 1),
(15, 1);`

	if _, err := conn.ExecContext(ctx, query3); err != nil {
		return err
	}

	query4 := `INSERT INTO orders_items (order_id, item_id, quantity) VALUES
(10, 1, 2),
(10, 3, 1),
(10, 6, 1),
(11, 2, 3),
(14, 1, 3),
(14, 4, 4),
(15, 5, 1);`

	if _, err := conn.ExecContext(ctx, query4); err != nil {
		return err
	}

	query5 := `INSERT INTO items_shelves (item_id, shelf_id) VALUES
(3, (SELECT id FROM shelves WHERE name = 'З')),
(3, (SELECT id FROM shelves WHERE name = 'В')),
(5, (SELECT id FROM shelves WHERE name = 'А'));`

	if _, err := conn.ExecContext(ctx, query5); err != nil {
		return err
	}

	return nil
}
