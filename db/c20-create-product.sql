CREATE TABLE IF NOT EXISTS products (
    product_id SERIAL PRIMARY KEY,
<<<<<<<< HEAD:db/c02-create-product.sql
    store_id SERIAL references stores(store_id),
========
    store_id SERIAL REFERENCES stores(store_id),
>>>>>>>> upstream/main:db/c20-create-product.sql
    name VARCHAR(255) NOT NULL,
    describe VARCHAR(255),
    category_id INT,
    picture BYTEA,
    price INT NOT NULL,
    stock INT NOT NULL
);
