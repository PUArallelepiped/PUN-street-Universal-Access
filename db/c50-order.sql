CREATE TABLE orders (
    cart_id INTEGER,
    store_id INTEGER REFERENCES stores(store_id),
    user_id INTEGER REFERENCES user_data(user_id),
    discount_id INTEGER REFERENCES discounts(discount_id),
    status INTEGER,
    total_price INTEGER,
    Order_date TIMESTAMP,
    taking_address VARCHAR(255),
    taking_method INTEGER,
    PRIMARY KEY (cart_id, store_id, user_id)
)