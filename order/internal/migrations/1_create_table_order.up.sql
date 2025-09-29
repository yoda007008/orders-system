CREATE TABLE IF NOT EXISTS "orders" (
    id SERIAL PRIMARY KEY,
    order_id INT,
    order_name TEXT,
    order_price INT,
    order_count INT,
    order_created_at TIMESTAMP,
    order_updated_at TIMESTAMP
);