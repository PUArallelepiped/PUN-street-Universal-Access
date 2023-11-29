CREATE TABLE IF NOT EXISTS carts (
    customer_id SERIAL REFERENCES user_data(user_id),
    product_id SERIAL REFERENCES products(product_id),
    store_id SERIAL REFERENCES stores(store_id),
    cart_id INTEGER NOT NULL,
    product_quantity INT NOT NULL,
    event_discount_id SERIAL REFERENCES discounts(discount_id),
    PRIMARY KEY (customer_id, product_id, store_id, cart_id)
);