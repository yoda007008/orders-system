CREATE TABLE IF NOT EXISTS order (
    id INT NOT NULL AUTO_INCREMENT,
    order_id INT NOT NULL,
    order_name TEXT NOT NULL,
    order_price INT NOT NULL,
    order_count INT NOT NULL,
    order_created_at DATETIME NOT NULL,
    order_updated_at DATETIME NOT NULL,
);