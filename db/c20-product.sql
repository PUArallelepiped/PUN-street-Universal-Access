CREATE TABLE IF NOT EXISTS products (
    product_id SERIAL PRIMARY KEY,
    store_id SERIAL REFERENCES stores(store_id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category_id INT,
    picture VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    stock INT NOT NULL
    status INT NOT NULL
);
