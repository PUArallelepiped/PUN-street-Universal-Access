CREATE TABLE
    IF NOT EXISTS event_discount (
        discount_id SERIAL UNIQUE REFERENCES discounts(discount_id),
        max_quantity INT NOT NULL,
        product_id SERIAL REFERENCES products(product_id)
    );