CREATE TABLE IF NOT EXISTS carts (
    customer_id SERIAL REFERENCES user_data(user_id),
    product_id SERIAL REFERENCES products(product_id),
    store_id SERIAL REFERENCES stores(store_id),
    cart_id SERIAL REFERENCES user_data(current_cart_id),
    price INT NOT NULL,
    product_quantity INT NOT NULL,
    discount_id SERIAL REFERENCES discount(discount_id),
    PRIMARY KEY (customer_id, product_id, store_id)
);