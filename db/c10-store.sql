CREATE TABLE IF NOT EXISTS stores (
    store_id SERIAL PRIMARY KEY REFERENCES user_data(user_id),
    name VARCHAR(255) NOT NULL,
    rate FLOAT NOT NULL,
    rate_count INTEGER NOT NULL,
    address VARCHAR(255) NOT NULL,
    picture VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    shipping_fee INTEGER NOT NULL,
    status INTEGER NOT NULL
);