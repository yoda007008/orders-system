CREATE TABLE IF NOT EXISTS "order" (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    order_name TEXT NOT NULL,
    order_price INT NOT NULL,
    order_count INT NOT NULL,
    order_created_at TIMESTAMP NOT NULL,
    order_updated_at TIMESTAMP NOT NULL
);