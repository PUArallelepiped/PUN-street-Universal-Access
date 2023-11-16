CREATE TABLE IF NOT EXISTS discounts (
    discount_id SERIAL PRIMARY KEY UNIQUE,
    discount_type INT NOT NULL,
    status INT NOT NULL,
    description TEXT,
    name VARCHAR(255) NOT NULL
);