CREATE TABLE IF NOT EXISTS items_shelves
(
    item_id    INT,
    shelf_id   INT,
    PRIMARY KEY (item_id, shelf_id),
    CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES items (id),
    CONSTRAINT fk_shelf FOREIGN KEY (shelf_id) REFERENCES shelves (id)
);