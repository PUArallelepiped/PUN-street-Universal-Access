CREATE TABLE IF NOT EXISTS carts (
    customer_id SERIAL REFERENCES user_datas(user_id),
    product_id SERIAL REFERENCES products(product_id),
    store_id SERIAL REFERENCES stores(store_id),
    product_quantity INT NOT NULL,
    PRIMARY KEY (customer_id, product_id, store_id)
);