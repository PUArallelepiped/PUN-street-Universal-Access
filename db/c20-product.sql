CREATE TABLE IF NOT EXISTS products (
    product_id SERIAL PRIMARY KEY,
    store_id SERIAL REFERENCES stores(store_id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    picture VARCHAR(255),
    price INT NOT NULL,
    stock INT NOT NULL,
    status INT NOT NULL
);
