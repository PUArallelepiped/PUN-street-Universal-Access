CREATE TABLE IF NOT EXISTS cart (
    customer_id SERIAL REFERENCES UserData(id),
    product_id SERIAL REFERENCES product(product_id),
    store_id SERIAL REFERENCES stores(store_id),
    product_quantity INT NOT NULL,
    PRIMARY KEY (customer_id, product_id, store_id)
);