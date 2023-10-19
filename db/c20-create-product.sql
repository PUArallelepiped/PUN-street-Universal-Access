CREATE TABLE IF NOT EXISTS products (
    product_id SERIAL PRIMARY KEY,
    store_id SERIAL REFERENCES stores(store_id),
    name VARCHAR(255) NOT NULL,
    describe VARCHAR(255),
    category_id INT,
    picture BYTEA,
    price INT NOT NULL,
    stock INT NOT NULL
);
