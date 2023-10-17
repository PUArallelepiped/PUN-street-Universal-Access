CREATE TABLE IF NOT EXISTS cart (
    customer_id SERIAL references UserData(id),
    product_id SERIAL references product(product_id),
    store_id SERIAL references stores(store_id),
    product_quantity INT NOT NULL,
    PRIMARY KEY (customer_id, product_id, store_id)
);