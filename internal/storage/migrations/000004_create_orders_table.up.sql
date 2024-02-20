CREATE TABLE IF NOT EXISTS orders
(
    id          SERIAL PRIMARY KEY,
    user_id     INT NOT NULL REFERENCES users (id)
);