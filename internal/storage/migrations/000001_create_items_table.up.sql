CREATE table IF NOT EXISTS items
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(255) NOT NULL,
    main_shelf_id  INT
);