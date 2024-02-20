CREATE TABLE IF NOT EXISTS orders_items
(
    order_id    INT,
    item_id   INT,
    quantity   INT,
    PRIMARY KEY (order_id, item_id),
    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (id),
    CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES items (id)
);